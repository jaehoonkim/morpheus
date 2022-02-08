package events_test

import (
	"log"
	"testing"

	"github.com/NexClipper/sudory/pkg/server/events"
	"github.com/NexClipper/sudory/pkg/server/macro/channels"
	"github.com/jinzhu/configor"
)

const test_config_filename = "test-event-config.yml"

type Configs struct {
	Events []Config `yaml:"events,omitempty"`
}

type Config struct {
	Name      string           `yaml:"name,omitempty"`
	Pattern   string           `yaml:"pattern,omitempty"`
	Listeners []ListenerConfig `yaml:"listeners,omitempty"`
}

type ListenerConfig struct {
	Type   string                 `yaml:"type,omitempty"`
	Option map[string]interface{} `yaml:"option,omitempty"`
}

func TestNewPasudoConfig(t *testing.T) {
	var err error

	cfg := Configs{}
	if err = configor.Load(&cfg, test_config_filename); err != nil {
		t.Fatal(err)
	}
}

func TestNewConfig(t *testing.T) {
	var err error
	var ecfg *events.Config
	var EventContexts []events.EventContexter

	if ecfg, err = events.NewConfig(test_config_filename); err != nil { //config file load
		t.Fatal(err)
	}
	if err = ecfg.Vaild(); err != nil { //config vaild
		t.Fatal(err)
	}
	if EventContexts, err = ecfg.MakeEventListener(); err != nil { //events regist listener
		t.Fatal(err)
	}

	for n, it := range ecfg.Events {
		t.Log(n)
		t.Log(it)
	}

	for n, it := range EventContexts {
		t.Log(n)
		t.Log(it)
	}
}

func TestActivate(t *testing.T) {
	var err error

	//events
	var eventContexts []events.EventContexter
	var eventConfig *events.Config
	//event config
	if eventConfig, err = events.NewConfig(test_config_filename); err != nil { //config file load
		panic(err)
	}
	//config vaild
	if err = eventConfig.Vaild(); err != nil {
		panic(err)
	}
	//event listener
	if eventContexts, err = eventConfig.MakeEventListener(); err != nil { //events regist listener
		panic(err)
	}
	//event manager
	eventInvoke := channels.NewSafeChannel(0)
	manager := events.NewManager(eventContexts, log.Printf)
	deactivate := manager.Activate(eventInvoke, len(eventContexts)) //manager activate
	defer func() {
		deactivate() //stop when closing
	}()
	events.Invoke = func(v *events.EventArgs) { eventInvoke.SafeSend(v) } //setting invoker

	count := 20

	for i := 0; i < count; i++ {
		events.Invoke(&events.EventArgs{
			Sender: "/client/auth",
			Args: map[string]interface{}{
				"hello": "workd",
				"count": i,
			},
		})
	}
}