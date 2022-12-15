package mygomod

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ====== serve.yaml sdk struct define ======
type SdkYaml struct {
	Basic          `yaml:",inline"`
	ModelTemplates []*ModelTemplate `yaml:"model_templates" json:"modelTemplates"`
	Inference      *SdkInference    `yaml:"inference" json:"inference"`
	Application    *SdkApplication  `yaml:"application" json:"application"`
	Single         *SdkSingleInfer  `yaml:"single" json:"single"`
}
type Basic struct {
	Type        string `yaml:"type" json:"type"`
	Vendor      string `yaml:"vendor" json:"vendor"`
	App         string `yaml:"app" json:"app"`
	Version     string `yaml:"version" json:"version"`
	Name        string `yaml:"name" json:"name"`
	Image       string `yaml:"image" json:"image"`
	Description string `yaml:"description" json:"description"`
}
type ModelTemplate struct {
	Title       string       `yaml:"title" json:"title"`
	Field       string       `yaml:"field" json:"field"`
	Task        string       `yaml:"task" json:"task"`
	SysParams   []*SysParam  `yaml:"sys_params" json:"sysParams"`
	UserParams  []*UserParam `yaml:"user_params" json:"userParams"`
	MaximumNum  int          `yaml:"maximum_num" json:"maximumNum"`
	MinimumNum  int          `yaml:"minimum_num" json:"minimumNum"`
	AvailModels []string     `yaml:"avail_models" json:"availModels"`
	//@add: 增加评判参数模板
	InspectParams []*UserParam `yaml:"inspect_params" json:"inspectParams"`
	Models        []*SdkModel  `json:"models"`
}
type SdkInference struct {
	Kernel       string    `yaml:"kernel" json:"kernel"`
	Agent        string    `yaml:"agent" json:"agent"`
	Passthrough  string    `yaml:"passthrough" json:"passthrough"`
	Proto        string    `yaml:"proto" json:"proto"`
	Engine       string    `yaml:"engine" json:"engine"`
	Entrypoint   string    `yaml:"entrypoint" json:"entrypoint"`
	HealthyDelay int       `yaml:"healthy_delay" json:"healthyDelay"`
	Devices      []*Device `yaml:"devices" json:"devices"`
}
type Device struct {
	Type      string  `yaml:"type" json:"type"`
	Series    string  `yaml:"series" json:"series"`
	DeviceNum int     `yaml:"device_num" json:"deviceNum"`
	CPU       int     `yaml:"cpu" json:"cpu"`
	Memory    float64 `yaml:"memory" json:"memory"`
}
type Inspect struct {
	Engine     string `yaml:"engine" json:"engine"`
	Entrypoint string `yaml:"entrypoint" json:"entrypoint"`
	//@mark: 评估参数模板在 model_templates中
}
type SdkApplication struct {
	//SceneTag   string      `yaml:"sceneTag" json:"sceneTag"`
	SysParams  []*SysParam  `yaml:"sys_params" json:"sysParams"`
	UserParams []*UserParam `yaml:"user_params" json:"userParams"`
	//@add: mark this app is a single model wrapped application
	Single bool `yaml:"single" json:"single"`
	//@add: mlops评判相关配置
	Inspect *Inspect `yaml:"inspect" json:"inspect"`
}
type SdkModel struct {
	Field     string `yaml:"field" json:"field"`
	Task      string `yaml:"task" json:"task"`
	LabId     int64  `json:"labId"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Id        int64  `json:"id"`
	VersionId int64  `json:"versionId"`
	Title     string `json:"title"`
	TitleTag  string `json:"titleTag"`
}
type SysParam struct {
	//平台约定的名称,应保持全局命名空间可见
	Name string `yaml:"name" json:"name"`
	//由平台在合适的时机填充
	Value string `yaml:"value" json:"value"`
	//Source string `yaml:"source" json:"source"`
	//传递给具体应用的名称,由app局部命名空间可见即可
	Arg string `yaml:"arg" json:"arg"`
	//Type     string `yaml:"type" json:"type"`
	//默认系统参数全部是必须的,除非optinal为true可以不传
	Optional bool `yaml:"optional" json:"optional"`
}
type UserParam struct {
	Name  string ` yaml:"name" json:"name"`
	Title string `yaml:"title" json:"title"`
	Type  string `yaml:"type" json:"type"`
	Value string `yaml:"value" json:"value"`
	//	Source   string `yaml:"source" json:"source"`
	Editable bool   `yaml:"editable" json:"editable"`
	Default  string `yaml:"default" json:"default"`
}
type SdkSingleInfer struct {
	PreProcess  *SdkSingleInferProcess `yaml:"preprocess" json:"preprocess"`
	PostProcess *SdkSingleInferProcess `yaml:"postprocess" json:"postprocess"`
	Infer       *SdkSingleNestedInfer  `yaml:"infer" json:"-"`
}
type SdkSingleInferProcess struct {
	Type       string                 `yaml:"type" json:"type"`
	CustomFile string                 `yaml:"custom_file" json:"custom_file"`
	ExtraFiles []interface{}          `yaml:"extra_files,flow" json:"extra_files"`
	Parameters []*SdkSingleInferParam `yaml:"parameters" json:"parameters"`
}
type SdkSingleInferParam struct {
	Name     string      `yaml:"name" json:"name"`
	Type     string      `yaml:"type" json:"type"`
	Value    interface{} `yaml:"value,flow" json:"value"`
	Mutable  bool        `yaml:"mutable" json:"mutable"`
	Editable bool        `yaml:"editable" json:"editable"`
}
type SdkSingleNestedInfer struct {
	Backend   string                       `yaml:"backend"`
	Device    string                       `yaml:"device"`
	DeviceIds string                       `yaml:"device_ids"`
	ModelFile string                       `yaml:"model_file"`
	Params    interface{}                  `yaml:"params"`
	Inputs    []*SdkSingleNestedInferParam `yaml:"inputs"`
	Outputs   []*SdkSingleNestedInferParam `yaml:"outputs"`
}

type SdkSingleNestedInferParam struct {
	Name  string      `yaml:"name"`
	Type  string      `yaml:"type"`
	Shape interface{} `yaml:"shape,flow"`
}

func yamlTest() {
	sdkYaml := SdkYaml{}
	// yamlFile := "/Users/jason/workspace/dummy/model_serve.yaml"
	yamlFile := "/Users/jason/workspace/apulis/project/iqi/sdk/gpu-det/cpp_serve.yaml"
	yamlContent, err := os.ReadFile(yamlFile)
	if err != nil {
		fmt.Println("read infer/transformer/serve.yaml", err)
		return
	}
	err = yaml.Unmarshal(yamlContent, &sdkYaml)
	if err != nil {
		fmt.Println("infer/transformer/serve.yaml unmarshal", err)
		return
	}
	sdkJson, err := json.Marshal(sdkYaml)
	if err != nil {
		fmt.Println("json.Marshal", err)
		return
	}
	fmt.Println(string(sdkJson))

	newYamlContent, err := yaml.Marshal(sdkYaml)
	if err != nil {
		fmt.Println("marshal newYamlContent", err)
		return
	}
	err = os.WriteFile("./new.yaml", newYamlContent, 0666)
	if err != nil {
		fmt.Println("write newYamlContent", err)
		return
	}
}
