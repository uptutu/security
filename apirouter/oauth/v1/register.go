/*
Copyright 2021 The tKeel Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"fmt"

	"github.com/tkeel-io/security/apirouter"
	"github.com/tkeel-io/security/apiserver/config"
	"github.com/tkeel-io/security/constants"
	"github.com/tkeel-io/security/models/oauth"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

func RegisterToRestContainer(c *restful.Container, conf *config.OAuth2Config) error {
	webservice := apirouter.GetWebserviceWithPatch(c, "/v1/oauth")

	oauthOperator, err := oauth.NewOperator(conf)
	if err != nil {
		_log.Error(err)
		return fmt.Errorf("oauth new operate %w", err)
	}
	handler := newOauthHandler(oauthOperator, conf)

	webservice.Route(webservice.GET("/authorize").
		To(handler.Authorize).
		Produces(restful.MIME_JSON).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagOauth}))

	webservice.Route(webservice.GET("/token").
		To(handler.Token).
		Doc("oauth token").
		Param(webservice.QueryParameter("grant_type", "GrantType:(password/authorization_code)").Required(true)).
		Param(webservice.QueryParameter("username", "user name while GrantType is password,style must be 'tenantID-username'")).
		Param(webservice.QueryParameter("password", "password")).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagOauth}))

	webservice.Route(webservice.GET("/authenticate").
		To(handler.Authenticate).
		Doc("oauth authenticate").
		Param(webservice.HeaderParameter("Authorization", "Bearer token")).
		Produces(restful.MIME_JSON).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagOauth}))

	webservice.Route(webservice.GET("/on_code").
		To(handler.OnCode).
		Produces(restful.MIME_JSON).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagOauth}))

	webservice.Route(webservice.GET("/{idp_type}/{tenant_id}/connection").
		To(handler.Connections).
		Doc("identity provider connection").
		Param(webservice.PathParameter("idp_type", "identity provider type: OIDC LDAP").Required(true)).
		Param(webservice.PathParameter("tenant_id", "tenant`s ID").Required(true)).
		Produces(restful.MIME_JSON).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagOauth}))

	webservice.Route(webservice.GET("/{idp_type}/{tenant_id}/callback").
		To(handler.OnConnections).
		Doc("identity provider connection callback address").
		Produces(restful.MIME_JSON).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagOauth}))

	webservice.Route(webservice.POST("/{idp_type}/{tenant_id}/register").
		To(handler.Register).
		Doc("register a external identity provider").
		Produces(restful.MIME_JSON).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.APITagOauth}))
	return nil
}
