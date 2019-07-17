package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
)

func GetSiteFromApi (siteId string) (*miapi.Site, *apierrors.ApiError){

	site := &miapi.Site{
		ID : siteId,
	}

	if err := site.Get(); err != nil {

		return nil,err
	}

	return site,nil

}
