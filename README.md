# 3141-operator

Kubernetes Operator running commands triggered by GCP PubSub Messages (a [3141-notification](https://github.com/300481/3141-notification))

## Configuration

### Operator

The Operator must be configured by environment variables.

The commands first arguments are **Repository URL**, **Commit ID**, **Pushed at time as Unix Epoch**.

The arguments specified with ARGS are attached.

Environment Variable|Description                                                          |Type
--------------------|---------------------------------------------------------------------|--------
`SYSTEM_ID`         |The ID of the system, let empty to get notifications for all systems.|*String*
`REF`               |The Git Repo Ref, let empty to get notifications for all refs.       |*String*
`COMMAND`           |The command to execute on notification.                              |*String*
`ARGS`              |The arguments for the command to execute.                            |*String*

### GCP PubSub

The GCP PubSub connector also must be configured by environment variables.

Please see [here](https://github.com/300481/mq) for this configuration.
