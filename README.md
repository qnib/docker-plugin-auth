# docker-plugin-authz
Docker authorisation plugin testbed

## Motivation

To make a read-only (only certain GET requests) unix socket, the authorization plugin needs to know by which unix-socket it was reached.
This plugin is an attempt to make this decission. Using the new plugin architecture, it's a bit uncovered ground yet. 
The current plugins use the `dockerd` option, which uses `github.com/docker/docker/pkg/authorization` and not `github.com/docker/go-plugins-helpers/sdk`.

For now the plugin starts but is not asked, when queried. :/

```bash
$ docker run -ti --rm --privileged --security-opt=seccomp=unconfined -v /:/host ubuntu tail -f /host/var/log/docker.log |grep plugin=
Aug 18 15:01:14 moby root: time="2017-08-18T15:01:14Z" level=info msg=">>>> Start plugin" plugin=def54c837ce8089588c4136b0153cd9c5b63334ba73292dad0c3e02fd96e63c9
```
