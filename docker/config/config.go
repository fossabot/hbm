package config

type Config struct {
	Action      string
	Description string
}

type Configs []Config

func New() Configs {
	config := []Config{}

	config = append(config, Config{Action: "container_create_param_privileged", Description: "--privileged param"})
	config = append(config, Config{Action: "container_create_param_ipc_host", Description: "--ipc=\"host\" param"})
	config = append(config, Config{Action: "container_create_param_net_host", Description: "--net=\"host\" param"})
	config = append(config, Config{Action: "container_create_param_pid_host", Description: "--pid=\"host\" param"})
	config = append(config, Config{Action: "container_create_param_userns_host", Description: "--userns=\"host\" param"})
	config = append(config, Config{Action: "container_create_param_uts_host", Description: "--uts=\"host\" param"})
	config = append(config, Config{Action: "container_create_param_user_root", Description: "--user=\"root\" param"})
	config = append(config, Config{Action: "container_create_param_publish_all", Description: "--publish-all param"})
	config = append(config, Config{Action: "container_create_param_security_opt", Description: "--security-opt param"})
	config = append(config, Config{Action: "container_create_param_sysctl", Description: "--sysctl param"})

	config = append(config, Config{Action: "image_create_official", Description: "Pull of Official image"})

	return config
}

func (c Configs) ActionExists(name string) bool {
	for _, config := range c {
		if config.Action == name {
			return true
		}
	}

	return false
}
