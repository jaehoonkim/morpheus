package control

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/NexClipper/sudory/pkg/server/control/vault"
	"github.com/NexClipper/sudory/pkg/server/macro/echoutil"
	"github.com/NexClipper/sudory/pkg/server/macro/logs"
	eventv1 "github.com/NexClipper/sudory/pkg/server/model/event/v1"
	metav1 "github.com/NexClipper/sudory/pkg/server/model/meta/v1"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

// Create Event
// @Description Create a event
// @Accept      json
// @Produce     json
// @Tags        server/event
// @Router      /server/event [post]
// @Param       x_auth_token header string          false "client session token"
// @Param       object       body   v1.Event_create true  "Event_create"
// @Success     200 {object} v1.EventWithEdges
func (ctl Control) CreateEvent(ctx echo.Context) error {

	body := new(eventv1.Event_create)
	if err := echoutil.Bind(ctx, body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(ErrorBindRequestObject(), "bind%s",
				logs.KVL(
					"type", TypeName(body),
				)))
	}

	if len(body.Name) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(ErrorInvalidRequestParameter(), "valid param%s",
				logs.KVL(
					ParamLog(fmt.Sprintf("%s.Name", TypeName(body)), body.Name)...,
				)))
	}

	//pattern
	if _, err := regexp.Compile(body.Pattern); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(err, "regexp compile event pattern"))
	}

	event := eventv1.Event{}
	event.UuidMeta = metav1.NewUuidMeta()
	event.LabelMeta = metav1.NewLabelMeta(body.Name, body.Summary)
	event.EventProperty = body.EventProperty

	r := eventv1.EventWithEdges{}
	_, err := ctl.ScopeSession(func(tx *xorm.Session) (interface{}, error) {
		//create event
		event_, err := vault.NewEvent(tx).Create(event)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
				errors.Wrapf(err, "create event"))
		}
		r.Event = *event_
		r.NotifierEdges = body.NotifierEdges

		//create notifier edges
		if err := AddEventNotifierEdges(tx, event_.Uuid, body.NotifierEdges); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
				errors.Wrapf(err, "create event notifier edge"))
		}

		return event_, err
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, r)
}

// Find Event
// @Description Find event
// @Accept      json
// @Produce     json
// @Tags        server/event
// @Router      /server/event [get]
// @Param       x_auth_token header string false "client session token"
// @Param       q            query  string false "query  pkg/server/database/prepared/README.md"
// @Param       o            query  string false "order  pkg/server/database/prepared/README.md"
// @Param       p            query  string false "paging pkg/server/database/prepared/README.md"
// @Success     200 {array} v1.Event
func (ctl Control) FindEvent(ctx echo.Context) error {
	//find event
	events, err := vault.NewEvent(ctl.db.Engine().NewSession()).Query(echoutil.QueryParam(ctx))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
			errors.Wrapf(err, "query event"))
	}

	return ctx.JSON(http.StatusOK, events)

}

// Get Event
// @Description Get a event
// @Accept      json
// @Produce     json
// @Tags        server/event
// @Router      /server/event/{uuid} [get]
// @Param       x_auth_token header string false "client session token"
// @Param       uuid         path   string true  "Event 의 Uuid"
// @Success     200 {object} v1.Event
func (ctl Control) GetEvent(ctx echo.Context) error {
	if len(echoutil.Param(ctx)[__UUID__]) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(ErrorInvalidRequestParameter(), "valid param%s",
				logs.KVL(
					ParamLog(__UUID__, echoutil.Param(ctx)[__UUID__])...,
				)))
	}

	uuid := echoutil.Param(ctx)[__UUID__]

	//get event
	event, err := vault.NewEvent(ctl.db.Engine().NewSession()).Get(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
			errors.Wrapf(err, "get event"))
	}

	return ctx.JSON(http.StatusOK, event)
}

// Get Event Edges
// @Description Get event edges
// @Accept      json
// @Produce     json
// @Tags        server/event
// @Router      /server/event/{uuid}/edges [get]
// @Param       x_auth_token header string false "client session token"
// @Param       uuid         path   string true  "Event 의 Uuid"
// @Success     200 {object} v1.EventNotifierEdge
func (ctl Control) GetEventEdges(ctx echo.Context) error {
	if len(echoutil.Param(ctx)[__UUID__]) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(ErrorInvalidRequestParameter(), "valid param%s",
				logs.KVL(
					ParamLog(__UUID__, echoutil.Param(ctx)[__UUID__])...,
				)))
	}

	uuid := echoutil.Param(ctx)[__UUID__]

	//get event
	event, err := vault.NewEvent(ctl.db.Engine().NewSession()).Get(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
			errors.Wrapf(err, "get event"))
	}

	//find edge
	edges, err := vault.NewEventNotifierEdge(ctl.db.Engine().NewSession()).Find("event_uuid = ?", event.Uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
			errors.Wrapf(err, "find edge"))
	}

	return ctx.JSON(http.StatusOK, edges)
}

// @Description Update a event
// @Accept      json
// @Produce     json
// @Tags        server/event
// @Router      /server/event/{uuid} [put]
// @Param       x_auth_token header string          false "client session token"
// @Param       uuid         path   string          true  "Event 의 Uuid"
// @Param       object       body   v1.Event_update true  "Event_update"
// @Success     200 {object} v1.Event
func (ctl Control) UpdateEvent(ctx echo.Context) error {
	body := new(eventv1.Event_update)
	if err := echoutil.Bind(ctx, body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(ErrorBindRequestObject(), "bind%s",
				logs.KVL(
					"type", TypeName(body),
				)))
	}

	uuid := echoutil.Param(ctx)[__UUID__]

	event := eventv1.Event{}
	event.Uuid = uuid
	event.LabelMeta = body.LabelMeta
	event.EventProperty = body.EventProperty

	r, err := ctl.ScopeSession(func(tx *xorm.Session) (interface{}, error) {
		event_, err := vault.NewEvent(tx).Update(event)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
				errors.Wrapf(err, "update event"))
		}

		// event = *event_
		return event_, nil
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, r)
}

// @Description addtion event notifier edge
// @Accept      json
// @Produce     json
// @Tags        server/event
// @Router      /server/event/{uuid}/edges/add [put]
// @Param       x_auth_token header string           false "client session token"
// @Param       uuid         path   string           true  "Event 의 Uuid"
// @Param       object       body   v1.NotifierEdges true  "NotifierEdges"
// @Success     200 {object} v1.EventNotifierEdge
func (ctl Control) UpdateEventAddtionNotifiers(ctx echo.Context) error {
	body := new(eventv1.NotifierEdges)
	if err := echoutil.Bind(ctx, body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(ErrorBindRequestObject(), "bind%s",
				logs.KVL(
					"type", TypeName(body),
				)))
	}

	uuid := echoutil.Param(ctx)[__UUID__]

	//get event
	event, err := vault.NewEvent(ctl.db.Engine().NewSession()).Get(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
			errors.Wrapf(err, "get event"))
	}

	//addtion event notifier edge
	_, err = ctl.ScopeSession(func(tx *xorm.Session) (interface{}, error) {
		if err := AddEventNotifierEdges(tx, event.Uuid, body.NotifierEdges); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
				errors.Wrapf(err, "addtion event notifier edge"))
		}
		return nil, nil
	})
	if err != nil {
		return err
	}

	edges, err := vault.NewEventNotifierEdge(ctl.db.Engine().NewSession()).Find("event_uuid = ?", uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
			errors.Wrapf(err, "find event edge"))
	}

	return ctx.JSON(http.StatusOK, edges)
}

// @Description subtraction event sub notifier
// @Accept      json
// @Produce     json
// @Tags        server/event
// @Router      /server/event/{uuid}/edges/sub [put]
// @Param       x_auth_token header string           false "client session token"
// @Param       uuid         path   string           true  "Event 의 Uuid"
// @Param       object       body   v1.NotifierEdges true  "NotifierEdges"
// @Success     200 {object} v1.EventNotifierEdge
func (ctl Control) UpdateEventSubtractionNotifiers(ctx echo.Context) error {
	body := new(eventv1.NotifierEdges)
	if err := echoutil.Bind(ctx, body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(ErrorBindRequestObject(), "bind%s",
				logs.KVL(
					"type", TypeName(body),
				)))
	}

	uuid := echoutil.Param(ctx)[__UUID__]

	//get event
	event, err := vault.NewEvent(ctl.db.Engine().NewSession()).Get(uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
			errors.Wrapf(err, "get event"))
	}

	//subtraction event sub notifier
	_, err = ctl.ScopeSession(func(tx *xorm.Session) (interface{}, error) {
		for _, edge := range body.NotifierEdges {
			if err := vault.NewEventNotifierEdge(tx).Delete(event.Uuid, edge.NotifierUuid); err != nil {
				return nil, echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
					errors.Wrapf(err, "subtraction event notifier edge"))
			}
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	edges, err := vault.NewEventNotifierEdge(ctl.db.Engine().NewSession()).Find("event_uuid = ?", uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
			errors.Wrapf(err, "find event edge"))
	}

	return ctx.JSON(http.StatusOK, edges)
}

// Delete Event
// @Description Delete a event
// @Accept json
// @Produce json
// @Tags server/event
// @Router /server/event/{uuid} [delete]
// @Param       x_auth_token header string false "client session token"
// @Param       uuid         path   string true  "Event 의 Uuid"
// @Success 200
func (ctl Control) DeleteEvent(ctx echo.Context) error {
	if len(echoutil.Param(ctx)[__UUID__]) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(
			errors.Wrapf(ErrorInvalidRequestParameter(), "valid param%s",
				logs.KVL(
					ParamLog(__UUID__, echoutil.Param(ctx)[__UUID__])...,
				)))
	}

	uuid := echoutil.Param(ctx)[__UUID__]

	_, err := ctl.ScopeSession(func(tx *xorm.Session) (interface{}, error) {
		//delete event
		if err := vault.NewEvent(tx).Delete(uuid); err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError).SetInternal(
				errors.Wrapf(err, "delete event"))
		}

		return nil, nil
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, OK())
}

func AddEventNotifierEdges(tx *xorm.Session, event_uuid string, edges []eventv1.NotifierEdge) error {
	for _, edge := range edges {
		//check notifier
		_, err := vault.NewEventNotifier(tx).Get(edge.NotifierType, edge.NotifierUuid)
		if err != nil {
			return errors.Wrapf(err, "get event notifier")
		}

		bind_edges, err := vault.NewEventNotifierEdge(tx).
			Find("event_uuid = ? AND notifier_type = ? AND notifier_uuid = ?",
				event_uuid, edge.NotifierType, edge.NotifierUuid)
		if err != nil {
			return errors.Wrapf(err, "find event notifier edge")
		}

		if 0 < len(bind_edges) {
			continue //already has
		}

		//create edge
		edge_ := eventv1.EventNotifierEdge{}
		edge_.EventUuid = event_uuid
		edge_.NotifierType = edge.NotifierType
		edge_.NotifierUuid = edge.NotifierUuid

		if _, err := vault.NewEventNotifierEdge(tx).Create(edge_); err != nil {
			return errors.Wrapf(err, "create event notifier edge")
		}
	}

	return nil
}