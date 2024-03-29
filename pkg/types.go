package pkg

import "time"

type DoryConfig struct {
	ServerURL   string        `yaml:"serverURL" json:"serverURL" bson:"serverURL" validate:""`
	Insecure    bool          `yaml:"insecure" json:"insecure" bson:"insecure" validate:""`
	Timeout     time.Duration `yaml:"timeout" json:"timeout" bson:"timeout" validate:""`
	AccessToken string        `yaml:"accessToken" json:"accessToken" bson:"accessToken" validate:""`
	UserToken   string        `yaml:"userToken" json:"userToken" bson:"userToken" validate:""`
}

type InstallDockerImage struct {
	Source string `yaml:"source" json:"source" bson:"source" validate:"required"`
	Target string `yaml:"target" json:"target" bson:"target" validate:""`
}

type InstallDockerImages struct {
	InstallDockerImages []InstallDockerImage `yaml:"dockerImages" json:"dockerImages" bson:"dockerImages" validate:""`
}

type InstallConfig struct {
	InstallMode string `yaml:"installMode" json:"installMode" bson:"installMode" validate:"required"`
	RootDir     string `yaml:"rootDir" json:"rootDir" bson:"rootDir" validate:"required"`
	HostIP      string `yaml:"hostIP" json:"hostIP" bson:"hostIP" validate:"required"`
	ViewURL     string `yaml:"viewURL" json:"viewURL" bson:"viewURL" validate:"required"`
	Dory        struct {
		Namespace    string            `yaml:"namespace" json:"namespace" bson:"namespace" validate:"required"`
		NodeSelector map[string]string `yaml:"nodeSelector" json:"nodeSelector" bson:"nodeSelector" validate:""`
		GitRepo      struct {
			Type    string `yaml:"type" json:"type" bson:"type" validate:"required"`
			Image   string `yaml:"image" json:"image" bson:"image" validate:"required"`
			ImageDB string `yaml:"imageDB" json:"imageDB" bson:"imageDB" validate:""`
			Port    int    `yaml:"port" json:"port" bson:"port" validate:"required"`
		} `yaml:"gitRepo" json:"gitRepo" bson:"gitRepo" validate:"required"`
		ArtifactRepo struct {
			Type     string `yaml:"type" json:"type" bson:"type" validate:"required"`
			Image    string `yaml:"image" json:"image" bson:"image" validate:"required"`
			Port     int    `yaml:"port" json:"port" bson:"port" validate:"required"`
			PortHub  int    `yaml:"portHub" json:"portHub" bson:"portHub" validate:"required"`
			PortGcr  int    `yaml:"portGcr" json:"portGcr" bson:"portGcr" validate:"required"`
			PortQuay int    `yaml:"portQuay" json:"portQuay" bson:"portQuay" validate:"required"`
		} `yaml:"artifactRepo" json:"artifactRepo" bson:"artifactRepo" validate:"required"`
		Openldap struct {
			Image      string `yaml:"image" json:"image" bson:"image" validate:"required"`
			ImageAdmin string `yaml:"imageAdmin" json:"imageAdmin" bson:"imageAdmin" validate:"required"`
			Port       int    `yaml:"port" json:"port" bson:"port" validate:"required"`
			Password   string `yaml:"password" json:"password" bson:"password" validate:""`
			Domain     string `yaml:"domain" json:"domain" bson:"domain" validate:"required"`
			BaseDN     string `yaml:"baseDN" json:"baseDN" bson:"baseDN" validate:"required"`
		} `yaml:"openldap" json:"openldap" bson:"openldap" validate:"required"`
		Redis struct {
			Image    string `yaml:"image" json:"image" bson:"image" validate:"required"`
			Password string `yaml:"password" json:"password" bson:"password" validate:""`
		} `yaml:"redis" json:"redis" bson:"redis" validate:"required"`
		Mongo struct {
			Image    string `yaml:"image" json:"image" bson:"image" validate:"required"`
			Password string `yaml:"password" json:"password" bson:"password" validate:""`
		} `yaml:"mongo" json:"mongo" bson:"mongo" validate:"required"`
		Docker struct {
			Image        string `yaml:"image" json:"image" bson:"image" validate:"required"`
			DockerName   string `yaml:"dockerName" json:"dockerName" bson:"dockerName" validate:"required"`
			DockerNumber int    `yaml:"dockerNumber" json:"dockerNumber" bson:"dockerNumber" validate:"required"`
		} `yaml:"docker" json:"docker" bson:"docker" validate:"required"`
		Dorycore struct {
			Port int `yaml:"port" json:"port" bson:"port" validate:"required"`
		} `yaml:"dorycore" json:"dorycore" bson:"dorycore" validate:"required"`
	} `yaml:"dory" json:"dory" bson:"dory" validate:"required"`
	ImageRepo struct {
		Namespace        string `yaml:"namespace" json:"namespace" bson:"namespace" validate:"required"`
		Type             string `yaml:"type" json:"type" bson:"type" validate:"required"`
		DomainName       string `yaml:"domainName" json:"domainName" bson:"domainName" validate:"required"`
		Version          string `yaml:"version" json:"version" bson:"version" validate:"required"`
		Password         string `yaml:"password" json:"password" bson:"password" validate:""`
		CertsDir         string `yaml:"certsDir" json:"certsDir" bson:"certsDir" validate:""`
		DataDir          string `yaml:"dataDir" json:"dataDir" bson:"dataDir" validate:""`
		RegistryPassword string `yaml:"registryPassword" json:"registryPassword" bson:"registryPassword" validate:""`
		RegistryHtpasswd string `yaml:"registryHtpasswd" json:"registryHtpasswd" bson:"registryHtpasswd" validate:""`
		VersionBig       string `yaml:"versionBig" json:"versionBig" bson:"versionBig" validate:""`
	} `yaml:"imageRepo" json:"imageRepo" bson:"imageRepo" validate:"required"`
	Dorycore struct {
		AdminUser struct {
			Username string `yaml:"username" json:"username" bson:"username" validate:"required"`
			Name     string `yaml:"name" json:"name" bson:"name" validate:"required"`
			Mail     string `yaml:"mail" json:"mail" bson:"mail" validate:"required"`
			Mobile   string `yaml:"mobile" json:"mobile" bson:"mobile" validate:"required"`
		} `yaml:"adminUser" json:"adminUser" bson:"adminUser" validate:"required"`
		Mail struct {
			Host     string `yaml:"host" json:"host" bson:"host" validate:"required"`
			Port     int    `yaml:"port" json:"port" bson:"port" validate:"required"`
			Username string `yaml:"username" json:"username" bson:"username" validate:"required"`
			Password string `yaml:"password" json:"password" bson:"password" validate:"required"`
			Ssl      bool   `yaml:"ssl" json:"ssl" bson:"ssl" validate:""`
			From     string `yaml:"from" json:"from" bson:"from" validate:"required"`
		} `yaml:"mail" json:"mail" bson:"mail" validate:"required"`
	} `yaml:"dorycore" json:"dorycore" bson:"dorycore" validate:"required"`
	Kubernetes struct {
		Host          string `yaml:"host" json:"host" bson:"host" validate:"required"`
		Port          int    `yaml:"port" json:"port" bson:"port" validate:"required"`
		Token         string `yaml:"token" json:"token" bson:"token" validate:"required"`
		PvConfigLocal struct {
			LocalPath string `yaml:"localPath" json:"localPath" bson:"localPath" validate:""`
		} `yaml:"pvConfigLocal" json:"pvConfigLocal" bson:"pvConfigLocal" validate:""`
		PvConfigNfs struct {
			NfsPath   string `yaml:"nfsPath" json:"nfsPath" bson:"nfsPath" validate:""`
			NfsServer string `yaml:"nfsServer" json:"nfsServer" bson:"nfsServer" validate:""`
		} `yaml:"pvConfigNfs" json:"pvConfigNfs" bson:"pvConfigNfs" validate:""`
		PvConfigCephfs struct {
			CephPath     string   `yaml:"cephPath" json:"cephPath" bson:"cephPath" validate:""`
			CephUser     string   `yaml:"cephUser" json:"cephUser" bson:"cephUser" validate:""`
			CephSecret   string   `yaml:"cephSecret" json:"cephSecret" bson:"cephSecret" validate:""`
			CephMonitors []string `yaml:"cephMonitors" json:"cephMonitors" bson:"cephMonitors" validate:""`
		} `yaml:"pvConfigCephfs" json:"pvConfigCephfs" bson:"pvConfigCephfs" validate:""`
	} `yaml:"kubernetes" json:"kubernetes" bson:"kubernetes" validate:"required"`
}

type KubePodState struct {
	Waiting struct {
		Reason string `yaml:"reason" json:"reason" bson:"reason" validate:""`
	} `yaml:"waiting" json:"waiting" bson:"waiting" validate:""`
	Running struct {
		StartedAt string `yaml:"startedAt" json:"startedAt" bson:"startedAt" validate:""`
	} `yaml:"running" json:"running" bson:"running" validate:""`
	Terminated struct {
		Reason   string `yaml:"reason" json:"reason" bson:"reason" validate:""`
		ExitCode int    `yaml:"exitCode" json:"exitCode" bson:"exitCode" validate:""`
		Signal   int    `yaml:"signal" json:"signal" bson:"signal" validate:""`
	} `yaml:"terminated" json:"terminated" bson:"terminated" validate:""`
}

type KubePod struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion" bson:"apiVersion" validate:"required"`
	Kind       string `yaml:"kind" json:"kind" bson:"kind" validate:"required"`
	MetaData   struct {
		Name              string            `yaml:"name" json:"name" bson:"name" validate:"required"`
		NameSpace         string            `yaml:"namespace" json:"namespace" bson:"namespace" validate:""`
		Labels            map[string]string `yaml:"labels" json:"labels" bson:"labels" validate:""`
		Annotations       map[string]string `yaml:"annotations" json:"annotations" bson:"annotations" validate:""`
		CreationTimestamp string            `yaml:"creationTimestamp" json:"creationTimestamp" bson:"creationTimestamp" validate:""`
		DeletionTimestamp string            `yaml:"deletionTimestamp" json:"deletionTimestamp" bson:"deletionTimestamp" validate:""`
		OwnerReferences   []struct {
			ApiVersion         string `yaml:"apiVersion" json:"apiVersion" bson:"apiVersion" validate:"required"`
			BlockOwnerDeletion bool   `yaml:"blockOwnerDeletion" json:"blockOwnerDeletion" bson:"blockOwnerDeletion" validate:""`
			Controller         bool   `yaml:"controller" json:"controller" bson:"controller" validate:""`
			Kind               string `yaml:"kind" json:"kind" bson:"kind" validate:""`
			Name               string `yaml:"name" json:"name" bson:"name" validate:""`
			Uid                string `yaml:"uid" json:"uid" bson:"uid" validate:""`
		} `yaml:"ownerReferences" json:"ownerReferences" bson:"ownerReferences" validate:""`
	} `yaml:"metadata" json:"metadata" bson:"metadata" validate:"required"`
	Spec struct {
		Containers []struct {
			Name  string `yaml:"name" json:"name" bson:"name" validate:""`
			Image string `yaml:"image" json:"image" bson:"image" validate:""`
		} `yaml:"containers" json:"containers" bson:"containers" validate:""`
	} `yaml:"spec" json:"spec" bson:"spec" validate:""`
	Status struct {
		Reason     string `yaml:"reason" json:"reason" bson:"reason" validate:""`
		Conditions []struct {
			Type   string `yaml:"type" json:"type" bson:"type" validate:""`
			Status string `yaml:"status" json:"status" bson:"status" validate:""`
		} `yaml:"conditions" json:"conditions" bson:"conditions" validate:""`
		ContainerStatuses []struct {
			Name         string       `yaml:"name" json:"name" bson:"name" validate:""`
			Ready        bool         `yaml:"ready" json:"ready" bson:"ready" validate:""`
			Started      bool         `yaml:"started" json:"started" bson:"started" validate:""`
			RestartCount int          `yaml:"restartCount" json:"restartCount" bson:"restartCount" validate:""`
			State        KubePodState `yaml:"state" json:"state" bson:"state" validate:""`
		} `yaml:"containerStatuses" json:"containerStatuses" bson:"containerStatuses" validate:""`
		InitContainerStatuses []struct {
			Name         string       `yaml:"name" json:"name" bson:"name" validate:""`
			Ready        bool         `yaml:"ready" json:"ready" bson:"ready" validate:""`
			Started      bool         `yaml:"started" json:"started" bson:"started" validate:""`
			RestartCount int          `yaml:"restartCount" json:"restartCount" bson:"restartCount" validate:""`
			State        KubePodState `yaml:"state" json:"state" bson:"state" validate:""`
		} `yaml:"initContainerStatuses" json:"initContainerStatuses" bson:"initContainerStatuses" validate:""`
		Phase     string    `yaml:"phase" json:"phase" bson:"phase" validate:""`
		PodIP     string    `yaml:"podIP" json:"podIP" bson:"podIP" validate:""`
		StartTime time.Time `yaml:"startTime" json:"startTime" bson:"startTime" validate:""`
	} `yaml:"status" json:"status" bson:"status" validate:""`
}

type KubePodList struct {
	Items []KubePod `yaml:"items" json:"items" bson:"items" validate:""`
}
