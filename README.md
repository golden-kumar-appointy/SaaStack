# SaaStack

## Communication between Two plugins

```mermaid
sequenceDiagram
  autonumber
  participant P as ProductPlugin
  participant C as Core
  participant U as UserInterfaceHandler
  participant E as EmailInterfaceHandler
  participant UP as ClerkPlugin
  participant EP as EmailSESPlugin

  Note left of P: {pluginId: Clerk, config: config, data: data}
  P ->> C : user.getPreference({data})
  C ->> U : {data}
  U ->> UP : {data}
  UP ->> U : {response}
  UP ->> C: {response}
  C ->> P : {response}

  P ->> C : Email.send({data})
  Note left of P: {pluginId: EmailSES, config: config, data: data}
  C ->> E : {data}
  E ->> EP : {data}
  EP ->> E : {response}
  E ->> C : {response}
  C ->> P : {reponse}

```
