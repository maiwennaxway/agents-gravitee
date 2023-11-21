// Modified after generation
package models

// VirtualHost Virtual host details. See also <a href=\"https://docs.gravitee.com/api-platform/fundamentals/virtual-host-property-reference\">Virtual host property reference</a>.
type VirtualHost struct {
	// List of host aliases.   A host alias provides the publicly visible DNS name of the virtual host on the Router and optionally includes the port number. The combination of host alias name and port number for the virtual host must be unique for all virtual hosts in the Edge installation. That means multiple virtual hosts can use the same port number if they have different host aliases.  You must create a DNS entry and CNAME record that matches the host alias, and the host alias must match the string that the client passes in the `Host` header.   The port number in `HostAliases` is optional. If you specify the port as part of the host alias, you must also specify the same port by using the `Port` element. Or, you can specify two `HostAliases` elements, one with the port number and one without.    You can have multiple `HostAlias` definitions in the same virtual host definition, corresponding to multiple DNS entries for the virtual host, but not for multiple ports. If you want multiple ports, create multiple virtual host definitions with different ports.   You can include the `*` wildcard character in the host alias. The `*` wildcard character can only be at the start (preceding the first `.`) of the host alias, and cannot be mixed with other characters. For example `*.example.com`. The TLS cert for the virtual host must have a matching wildcard in the CN name of the cert. For example, `*.example.com`. Using a wildcard in a virtual host alias lets API proxies handle calls addressed to multiple subdomains such as `alpha.example.com`, `beta.example.com`, or `live.example.com`. Using a wildcard alias also helps you use fewer virtual hosts per environment to stay within product limits, since a virtual host with a wild card counts as only one virtual host.  **For Cloud**: If you have an existing virtual host that uses a port other than `443`, you cannot add or remove a host alias.   **For Private Cloud**: If you are setting the host alias by using the IP addresses of your Routers, and not DNS entries, add a separate host alias for each Router, specifying the IP address of each Router and port of the virtual host.
	HostAliases []string `json:"hostAliases,omitempty"`
	// **Edge for Private Cloud only.**   List of network interfaces that you want the `port` to be bound to. If you omit this element, the `port` is bound on all interfaces.
	Interfaces []string `json:"interfaces,omitempty"`
	// **Private Cloud 4.18.01 and later and for Edge Cloud by contacting <a href=\"https://cloud.google.com/gravitee/support\">gravitee Support</a>.**    If you use an ELB in TCP pass-thru mode to handle requests to the Edge Routers, the Router treats the IP address of the ELB as the client IP instead of the actual client IP. If the Router requires the true client IP, enable `proxy_protocol` on the ELB so that it passes the client IP in the TCP packet. On the Router, you must also set the `listenOption` on the virtual host to `proxy_protocol`. Because the ELB is in TCP pass-thru mode, you typically terminate TLS on the Router. Therefore, you usually only configure the virtual host to use proxy_protocol when you also configure it to use TLS.                     The default value for `listenOptions` is an empty string. To later unset `listenOptions`, update the virtual host and omit the `listenOptions` property from the payload.
	ListenOptions []string `json:"listenOptions,omitempty"`
	// Virtual host name. Valid values include: `A-Z0-9._\\-$%`
	Name string `json:"name,omitempty"`
	// Port number used by the virtual host. Ensure that the port is open on the Edge Router.  If you specify a port in a `hostAliases` element, then the port number specified by `port` must match it.  **For Cloud**: You must specify port `443` when creating a virtual host. If omitted, by default the port is set to `443`. If you have an existing virtual host that uses a port other than `443`, you cannot change the port.  **For Private Cloud releases 4.16.01 through 4.17.05**: When creating a virtual host, you specify the Router port used by the virtual host. For example, port `9001`. By default, the Router runs as the user `gravitee` which does not have access to privileged ports, typically ports `1024` and below. If you want to create a virtual host that binds the Router to a protected port then you have to configure the Router to run as a user with access to those ports. See <a href=\"https://docs.gravitee.com/private-cloud/latest/setting-virtual-host\">Setting up a virtual host</a> for more.  **For Private Cloud releases prior to 4.16.01**: A Router can listen to only one HTTPS connection per virtual host, on a specific port, with the specified cert. Therefore, multiple virtual hosts cannot use the same port number if TLS termination occurs on the Router at the specified port.
	Port string `json:"port,omitempty"`
	// Base URL that overrides the URL displayed by the Edge UI for an API proxy deployed to the virtual host. Useful when you have an external load balancer in front of the Edge Routers. See <a href=\"https://docs.gravitee.com/api-platform/system-administration/creating-virtual-host\">Configuring TLS access to an API for the Private Cloud</a> for more.  The value of BaseUrl must include the protocol (that is, `http://` or `https://`).
	BaseUrl string `json:"baseUrl,omitempty"`
	// Flag that specifies whether the OCSP (Online Certificate Status Protocol) client is enabled. The OSCP sends a status request to an OCSP responder to determine if the TLS certificate is valid. The response indicates if the TLS certificate is valid and not revoked.  When enabled, OCSP stapling allows Edge, acting as the TLS server for one-way TLS, to query the OCSP responder directly and then cache the response. Edge then returns this response to the TLS client, or staples it, as part of TLS handshaking. See <a href=\"https://www.digicert.com/enabling-ocsp-stapling.htm\">Enable OCSP Stapling on Your Server</a> for more.     TLS must be enabled to enable OCSP stapling. Set to `on` to enable. Defaults to `off`.
	OCSPStapling string `json:"oCSPStapling,omitempty"`
	// **Edge for Public Cloud and Edge for Private Cloud 4.18.01 and later**. Configuration that determines how the Router reacts for this virtual host when the Message Processor goes down.                   You can specify multiple values. Valid values include:                   * `off`: Disables retry and the virtual host returns a failure code upon a request.      * `http_599`(Default): If the Router receives an `HTTP 599` response from the Message Processor, the Router forwards the request to the next Message Processor. `HTTP 599` is a special response code that is generated by a Message Processor when it is being shut down. The Message Processor tries to complete all existing requests, but for any new requests it responds with `HTTP 599` to signal to the Router to retry the request on the next Message Processor.      * `error`: If an error occurred while establishing a connection with the Message Processor, passing a request to it, or reading the response header from it, the Router forwards the request to the next Message Processor.      * `timeout`:  If a timeout occurs while establishing a connection with the Message Processor, passing a request to it, or reading the response header from it, the Router forwards the request to the next Message Processor.      * `invalid_header`:  If the Message Processor returned an empty or invalid response, the Router forwards the request to the next Message Processor.      * `http_XXX`:  If the Message Processor returned a response with HTTP code `XXX`, the Router forwards the request to the next Message Processor.                   If you specify multiple values, the Router uses a logical OR to combine them.
	RetryOptions []string `json:"retryOptions,omitempty"`
	// SSL information.
	SSLInfo *SslInfo `json:"sSLInfo,omitempty"`
	// **Edge for Public Cloud only.**  Flag that specifies whether to use the gravitee freetrial cert and key. If you have a paid Edge for Cloud account and do not yet have a TLS cert and key, you can create a virtual host that uses the gravitee freetrial cert and key. That means you can create the virtual host without first creating a keystore.  The gravitee freetrial cert is defined for a domain of `*.gravitee.net`. Therefore, the `HostAlias` of the virtual host must also be in the form `*.gravitee.net`.  See <a href=\"https://docs.gravitee.com/api-platform/fundamentals/configuring-virtual-hosts-cloud#creatingavirtualhost\">Defining a virtual host that uses the gravitee freetrial cert and key</a>.
	UseBuiltInFreeTrailCert string                             `json:"useBuiltInFreeTrailCert,omitempty"`
	PropagateTLSInformation VirtualHostPropagateTlsInformation `json:"propagateTLSInformation,omitempty"`
	Properties              VirtualHostProperties              `json:"properties,omitempty"`
}
