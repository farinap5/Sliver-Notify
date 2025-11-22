Sliver Notify

This project aims to create a client, connectable with default operator protocol, to listen for **initial beacon** and **initial session** events and send notifications using default webhooks.

Run:
```bash
go run cmd/main.go config.yaml
```

Create a configuration file like the following:

If you don't wanna use Teams, just skip creating its configuration. Same for the other chats.

```yaml
path: "/tmp/test_0.0.0.0.cfg" # operator configuration file

telegram:
  token: "616468...:AAEr2okC..."
  chat: "-1001..."

discord:
  token: "yJOBLt2..."
  chat: "115..."

teams:
  webhook: "https://..."
```

![](/img/img1.png)