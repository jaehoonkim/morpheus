package v1

import (
	"testing"

	clusterv1 "github.com/NexClipper/sudory/pkg/server/model/cluster/v1"
	clsttknv1 "github.com/NexClipper/sudory/pkg/server/model/cluster_token/v1"
	eventv1 "github.com/NexClipper/sudory/pkg/server/model/event/v1"
	globvarv1 "github.com/NexClipper/sudory/pkg/server/model/global_variant/v1"
	servicev1 "github.com/NexClipper/sudory/pkg/server/model/service/v1"
	stepv1 "github.com/NexClipper/sudory/pkg/server/model/service_step/v1"
	sessionv1 "github.com/NexClipper/sudory/pkg/server/model/session/v1"
	templatev1 "github.com/NexClipper/sudory/pkg/server/model/template/v1"
	commandv1 "github.com/NexClipper/sudory/pkg/server/model/template_command/v1"
	trecipev1 "github.com/NexClipper/sudory/pkg/server/model/template_recipe/v1"
)

func TestSync(t *testing.T) {
	tests := []struct {
		name    string
		args    interface{}
		want    error
		wantErr bool
	}{
		{name: "clusterv1",
			args: new(clusterv1.Cluster), want: nil, wantErr: false},
		{name: "envv1",
			args: new(globvarv1.GlobalVariant), want: nil, wantErr: false},
		{name: "stepv1",
			args: new(stepv1.ServiceStep), want: nil, wantErr: false},
		{name: "servicev1",
			args: new(servicev1.Service), want: nil, wantErr: false},
		{name: "sessionv1",
			args: new(sessionv1.Session), want: nil, wantErr: false},
		{name: "commandv1",
			args: new(commandv1.TemplateCommand), want: nil, wantErr: false},
		{name: "templatev1",
			args: new(templatev1.Template), want: nil, wantErr: false},
		{name: "trecipev1",
			args: new(trecipev1.TemplateRecipe), want: nil, wantErr: false},
		{name: "tokenv1",
			args: new(clsttknv1.ClusterToken), want: nil, wantErr: false},
		{name: "eventv1:event",
			args: new(eventv1.Event), want: nil, wantErr: false},
		{name: "eventv1::notifier_edge",
			args: new(eventv1.EventNotifierEdge), want: nil, wantErr: false},
		{name: "eventv1::console",
			args: new(eventv1.EventNotifierConsole), want: nil, wantErr: false},
		{name: "eventv1::webhook",
			args: new(eventv1.EventNotifierWebhook), want: nil, wantErr: false},
		{name: "eventv1::rabbitMq",
			args: new(eventv1.EventNotifierRabbitMq), want: nil, wantErr: false},
		{name: "eventv1::notifier_status",
			args: new(eventv1.EventNotifierStatus), want: nil, wantErr: false},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println(tt.name)

			err := newEngine().Sync(tt.args)
			if (err != nil) && tt.wantErr {
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCondition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}