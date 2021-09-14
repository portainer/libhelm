package options

type InstallOptions struct {
	Name                    string
	Chart                   string
	Namespace               string
	Repo                    string
	Wait                    bool
	ValuesFile              string
	KubernetesClusterAccess *KubernetesClusterAccess
}
