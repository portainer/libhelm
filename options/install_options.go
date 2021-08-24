package options

type InstallOptions struct {
	Name                    string
	Chart                   string
	Namespace               string
	Repo                    string
	ValuesFile              string
	KubernetesClusterAccess *KubernetesClusterAccess
}
