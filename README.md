# Go-Podman

> A lightweight, opinionated, and no-BS Go client for communicating with [Podman](https://podman.io) via Unix socket.

## What Is This

`go-podman` is *not* trying to reinvent the wheel or emulate the Docker SDK.

It's a focused, modular interface to **interact with Podman's HTTP API** using Go - and nothing more.

We start with basic interactions, like managing containers or images, and leave room for future expansion. Need to add a new endpoint? Just build it.

## Key Features

- **Minimal abstraction**: Wraps Podman's REST API into modular, reusable Go components.
- **Unix socket communication**: Supports rootless and rootfull communication through Unix sockets.
- **Composability**: Build custom solutions by combining simple building blocks.
- **Incremental**: Only the endpoints you need — no huge SDK, no unnecessary overhead.

The goal is simple: **give you control over Podman via Go**, without forcing you to adopt unnecessary complexity.

## How It Works

We wrap Podman API requests in small, testable units. The library is divided into layers, including the **request building**, **response reading**, and **payload serialization**, with each of these easily replaceable for testing or customization.

Right now, we offer basic Podman interaction, and you're welcome to extend it as you need.

## Documentations

This project is modular and designed to evolve, so if you're looking to dig deeper into how it works or expand it for your needs, check out our detailed docs:

[Docs/What_Is_Go_Podman](https://github.com/MineHosting/go-podman/blob/main/docs/What_Is_Go_Podman.md)

## Contributing

If you're interested in adding support for new endpoints or improving the current architecture, feel free to contribute! We're keeping things simple and modular to let you extend the library with minimal friction.
