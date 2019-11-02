package youtube_api

type Params struct {
  Params []Param
}

type Param struct {
  key   string
  value string
}

func (params *Params) AddParam(key string, value string) *Params {
  param := Param{key, value}
  (*params).Params = append((*params).Params, param)
  return params
}

func (params *Params) GetParams() []Param {
  return (*params).Params
}

