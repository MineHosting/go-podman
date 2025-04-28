# Go Podman

**Go-Podman** is a *programmatic container manager* written in Go, designed to interact directly with the Podman API via Unix sockets.

It embraces the Unix Philosophy: *do one thing and do it well* - enabling you to build precise, composable tooling around container management.

---

## What It Does

Go Podman acts as a minimal abstraction layer over Podman's REST API, allowing developers to:

- **Manage containers in code**, without relying on CLI tools or shell scripts.
- **Automate complex workflows**, turning infrastructure logic into reproducible and testable Go functions.
- **Build custom orchestrators**, ephemeral enviroments, or internal tooling with fine-grained control.

---

## Philosophy

This project is guided by the following principles:

- **Simplicity** - Keep it small, composable, and easy to understand.
- **Modularity** - Each component is isolated and follow the Single Responsibility Principle
- **Extensibility** -  You can use individual building blocks or combine them for more complex solutions

---

## Use Cases

- Create a custom daemon that spins up containers on the fly
- Build internal devops tools (e.g CI/CD runners, sandbox enviroments)
- Interact with the podman Api with typed structures - no manual json parsing
- Replace brittle shell scripts with go-based, unit-testable logic

- **Reactive infrastructure scenarios**
Trigger container lifecycles in response to events - such as webhooks, queue messages ...
Perfect for building lightweight Function-as-a-Service (FaaS) systems, dev enviroments on-demand, or self-healing services without Kubernetes overhead.

## Disclaimer

This library is not a full-replacement for container orchestrators platforms like Kubernetes - it's low-level, flexible, local toolkit for building your own workflow on top of Podman.

If your team is allergic to [YAML](https://yaml.org/) bloat and wants something programmable, lightweight, and Go-native: you're in the right place
