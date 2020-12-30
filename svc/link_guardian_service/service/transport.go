package service

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"gitlab.com/vivsoft-platform/k8s-serverless/svc_kit/svc_pkg"
)

type link struct {
	Url         string
	Title       string
	Description string
	Status      string
	Tags        map[string]bool
	CreatedAt   string
	UpdatedAt   string
}

func newLink(source svc_pkg.Link) link {
	return link{
		Url:         source.Url,
		Title:       source.Title,
		Description: source.Description,
		Status:      source.Status,
		Tags:        source.Tags,
		CreatedAt:   source.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   source.UpdatedAt.Format(time.RFC3339),
	}
}

type getLinksResponse struct {
	Links []link `json:"links"`
	Err   string `json:"err"`
}

type deleteLinkRequest struct {
	Username string
	Url      string
}

type SimpleResponse struct {
	Err string `json:"err"`
}

func decodeGetLinksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request svc_pkg.GetLinksRequest
	q := r.URL.Query()
	request.UrlRegex = q.Get("url")
	request.TitleRegex = q.Get("title")
	request.DescriptionRegex = q.Get("description")
	request.Username = q.Get("username")
	request.Tag = q.Get("tag")
	request.StartToken = q.Get("start")
	return request, nil
}

func decodeAddLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request svc_pkg.AddLinkRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request svc_pkg.UpdateLinkRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteLinkRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteLinkRequest
	q := r.URL.Query()
	request.Username = q.Get("username")
	request.Url = q.Get("url")
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeGetLinksEndpoint(svc svc_pkg.LinkManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(svc_pkg.GetLinksRequest)
		result, err := svc.GetLinks(req)
		res := getLinksResponse{}
		for _, link := range result.Links {
			res.Links = append(res.Links, newLink(link))
		}
		if err != nil {
			res.Err = err.Error()
			return res, err
		}
		return res, nil
	}
}

func makeAddLinkEndpoint(svc svc_pkg.LinkManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(svc_pkg.AddLinkRequest)
		err := svc.AddLink(req)
		res := SimpleResponse{}
		if err != nil {
			res.Err = err.Error()
			return res, err
		}
		return res, nil
	}
}

func makeUpdateLinkEndpoint(svc svc_pkg.LinkManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(svc_pkg.UpdateLinkRequest)
		err := svc.UpdateLink(req)
		res := SimpleResponse{}
		if err != nil {
			res.Err = err.Error()
			return res, err
		}
		return res, nil
	}
}

func makeDeleteLinkEndpoint(svc svc_pkg.LinkManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteLinkRequest)
		err := svc.DeleteLink(req.Username, req.Url)
		res := SimpleResponse{}
		if err != nil {
			res.Err = err.Error()
			return res, err
		}
		return res, nil
	}
}