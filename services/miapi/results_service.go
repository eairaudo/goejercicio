package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
	"sync"
)

var wg sync.WaitGroup

var user *miapi.User

var country *miapi.Country

var site *miapi.Site

func GetResultFromApi(userId int64) (*miapi.Result, *apierrors.ApiError) {

	wg.Add(1)

	go GetUser(userId)

	wg.Wait()

	if user != nil {
		wg.Add(2)
		go GetSite(user.SiteID)
		go GetCountry(user.CountryID)
	}

	wg.Wait()

	result := &miapi.Result{
		User:    user,
		Site:    site,
		Country: country,
	}

	return result, nil

}

func GetUser(userId int64) *apierrors.ApiError {

	defer wg.Done()
	var err *apierrors.ApiError

	user, err = GetUserFromApi(userId)
	if err != nil {

		return err
	}

	return nil
}

func GetCountry(countryID string) *apierrors.ApiError {

	defer wg.Done()

	var err *apierrors.ApiError

	country, err = GetCountryFromApi(countryID)
	if err != nil {

		return err
	}

	return nil
}

func GetSite(siteID string) *apierrors.ApiError {

	defer wg.Done()

	var err *apierrors.ApiError

	site, err = GetSiteFromApi(siteID)
	if err != nil {

		return err
	}

	return nil
}
