You can see there is Claude.md file, and you might think I used claude to build this. 

See I used claude just to clarify few things, since i'm a newbie. you can read that md file. it says do not write any code, just give advice when I'm stuck somewhere.

## Checklist

- [x] TCP server accepting multiple concurrent client connections (one goroutine per connection)
- [x] Central hub managing clients safely via channels (register/unregister/broadcast)
- [x] Broadcast messaging — a message from one client is fanned out to all connected clients
- [x] Clients identified by remote address (`ip:port`) in broadcast messages
- [x] `QUIT` command to let a client disconnect cleanly
- [x] Basic CLI client with concurrent read (goroutine) and write (stdin loop)
- [x] client uses unique name instead of remote addr (name is given by client)
- [ ] private chat with a particular client (`/chat username` for private, `/chat broadcast` for public)
- [ ] join/leave notifications — broadcast a message to everyone when a client connects or disconnects, instead of them just silently appearing/vanishing
