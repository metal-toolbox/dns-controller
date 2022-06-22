# 2782 Manager

`2782-manager` is a simple controller that allows you to create and update [SRV records](https://datatracker.ietf.org/doc/html/rfc2782), as well as remove them when no endpoints are present.

## About

SRV records are a lightweight service discovery method that can be implemented in a wide range of applications. It uses the well-known DNS protocol to retrieve information about the location, port, and protocol of particular service. `2782-manager` centrally manages the lifecycle of these records.

The initial work will cover a server component and kubernetes controller that watching `Networking/v1 Ingress` objects. Inspired by [external-dns](https://github.com/kubernetes-sigs/external-dns), this is meant to work with an upstream DNS provider, like [ns1](ns1.com).

### Goals

For the server we hope to:

* [ ] Provide central service for creating and manage SRV records
* [ ] Create zones if they do not exist, delete them when no more endpoints are listed
* [ ] Test connectivity to endpoints and remove them if connections cannot be established

For the kubernetes controller:

* [ ] A kubernetes controller that watches `Networking/v1 Ingress` objects and has cluster local endpoints added or removed from the specified SRV record

## Non-goals

* Create a CRD for managing DNS. `external-dns` does this - well, what it does not do is provide a way to reconcile multiple clients making updates to a single record.
* Implement all possible DNS API's, if the project gets traction and a individuals want more integrations, open an issue and let's chat!

## Design

Follow the format `_service._proto.name. ttl IN SRV priority weight port target.`, the service will take a payload like

```go
type Answer struct {
  owner    string
  protocol string
  service  string
  target   string

  priority int
  port     int
  ttl      int
  weight   int
}
```

This would create a record like

```
_$service._$protocol.$owner.example.com $ttl IN SRV $weight $priority $port $target
```

So using this as a payload

```yaml
answer:
  owner: team-a
  priority: 0
  protocol: tcp
  port: 443
  service: artifacts
  target: artifacts.us1.example.com
  ttl: 360
  weight: 0
```

Would create a record like

```text
_artifacts._tcp.example.com 360 IN SRV 0 0 443 artifacts.us1.example.com
```
