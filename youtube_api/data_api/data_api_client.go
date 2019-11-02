package data_api

import (
  "conf"
)

type DataApiClient struct {
  Config conf.Config
}

func (dataApiClient *DataApiClient) Search() Search {
  return Search{
    (*dataApiClient).Config,
  }
}

func (dataApiClient *DataApiClient) Videos() Videos {
  return Videos{
    (*dataApiClient).Config,
  }
}
