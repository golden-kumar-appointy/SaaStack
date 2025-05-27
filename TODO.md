Step 1:

- Register 1 Interface
- Register 1 Plugin that implements the interface
- Make an HTTP API call to the platform which should resolve in the plugin
  - Core - Doesn’t Import
  - Interface - Can Import core only
  - Plugin - Can Import core and interface
  - Main - Can imports all three

Config:

```yaml
plugins:
  - name: clerk-auth
    source: grpc://my-custom-plugin.com:50051
    interface: auth
    dependencies: notification
    config:
      username: $user$
      password:
  - name:
```

Step 2:

### Todo

Lifecycle:

Lifecycle in uber fx

add database in plugin (for loging)

How to logging, which layer used logging.

Implement onion architechture, add logger,
chaning of one inside another

Database: for learning query

fx -> use life cycle to append user data.

after sending destroy the life cycle objects
