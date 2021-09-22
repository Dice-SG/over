package chargeratesort

import (
	"context"
	"os"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

// NodeName is a plugin that checks if a pod spec node name matches the current node.
type ChargeRateSort struct{}
type Nodes []node

var _ framework.FilterPlugin = &ChargeRateSort{}

type node struct {
	name       string
	chargerate int
}

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "ChargeRateSort"

	// ErrReason returned when node name doesn't match.
	ErrReason = "node(s) didn't match the requested node name"
)

// Name returns name of the plugin. It is used in logs, etc.
func (pl *ChargeRateSort) Name() string {
	return Name
}

// Filter invoked at the filter extension point.
func (pl *ChargeRateSort) Filter(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	if nodeInfo.Node() == nil {
		return framework.NewStatus(framework.Error, "node not found")
	}
	if !Fits(nodeInfo) {
		return framework.NewStatus(framework.UnschedulableAndUnresolvable, ErrReason)
	}
	return nil
}

// New initializes a new plugin and returns it.
func New(_ runtime.Object, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &ChargeRateSort{}, nil
}

/*
func Sort() string {
	node1 := node{"node1", 70}
	node2 := node{"node2", 30}
	node3 := node{"node3", 90}

	nodes := []node{node1, node2, node3}

	tmp := node{" ", 0}

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].chargerate < nodes[j].chargerate {
				tmp = nodes[i]
				nodes[i] = nodes[j]
				nodes[j] = tmp
			}
		}
	}

	return nodes[0].name
}

// Fits actually checks if the pod fits the node.
func Fits(name  , nodeInfo *framework.NodeInfo) bool {
	return name == nodeInfo.Node().Name
}
*/

func Fits(nodeInfo *framework.NodeInfo) bool {
	node1 := node{"node1", 70}
	node2 := node{"node2", 30}
	node3 := node{"node3", 90}

	nodes := []node{node1, node2, node3}

	tmp := node{" ", 0}

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].chargerate < nodes[j].chargerate {
				tmp = nodes[i]
				nodes[i] = nodes[j]
				nodes[j] = tmp
			}
		}
	}

	return nodes[0].name == nodeInfo.Node().Name
}

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(Name, New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
