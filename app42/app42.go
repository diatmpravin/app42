package app42

//http://play.golang.org/p/FVi-b1mywd

type Framework struct {
	Response struct {
		Success    bool
		Frameworks []struct {
			Id      string
			Name    string
			Version string
		}
	}
}

type Appframeworks struct {
	App42 Framework
}

type Runtime struct {
	Response struct {
		Success  bool
		Runtimes []struct {
			Id      string
			Name    string
			Version string
		}
	}
}

type AppRuntimes struct {
	App42 Runtime
}

type IaaS struct {
	Response struct {
		Success bool
		Iaas    []struct {
			Id   string
			Name string
			Zone string
		}
	}
}

type IaaSProviders struct {
	App42 IaaS
}

type Subscription struct {
	Response struct {
		Success        bool
		DeploymentType []string
	}
}

type AppSubscription struct {
	App42 Subscription
}

type ResponseData struct {
	Response struct {
		UserId       string `json:"userId"`
		Success      bool   `json:"success"`
		ResourceName string `json:"resourceName"`
		Description  string `json:"description"`
	}
}

type AppAvailability struct {
	App42 ResponseData
}

type App struct {
	Response struct {
		Success bool
		Apps    []struct {
			AppUrl         string
			AppStatus      string
			IaasProvider   string
			Email          string
			VmType         string
			AppState       string
			Name           string
			IaasIdentifier string
			Runtime        string
			InstanceCount  int
			Framework      string
			Memory         string
		}
	}
}

type AllApps struct {
	App42 App
}
