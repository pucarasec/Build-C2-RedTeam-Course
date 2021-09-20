import asyncio
import argparse


class Proxy:
    def __init__(self, bind_host: str, bind_port: int, target_host: str, target_port: int):
        self._bind_host = bind_host
        self._bind_port = bind_port
        self._target_host = target_host
        self._target_port = target_port

    async def forward_endpoint_data(self, reader, writer):
        while True:
            data = await reader.read(1024)
            if len(data) == 0: break
            writer.write(data)
            await writer.drain()


    async def handle_connection(self, client_reader, client_writer):
        print('New connection')
        server_reader, server_writer = await asyncio.open_connection(self._target_host, self._target_port)

        asyncio.create_task(self.forward_endpoint_data(client_reader, server_writer))
        asyncio.create_task(self.forward_endpoint_data(server_reader, client_writer))

    async def run(self):
        server = await asyncio.start_server(self.handle_connection,
                                            self._bind_host, self._bind_port)

        addr = server.sockets[0].getsockname()

        async with server:
            await server.serve_forever()

def main():
    parser = argparse.ArgumentParser(description='TCP Proxy')
    parser.add_argument('--bind_host', '-b', type=str, required=False, default='127.0.0.1', help='Bind hostname')
    parser.add_argument('bind_port', type=int, help='Bind port')
    parser.add_argument('target_host', type=str, help='Target hostname')
    parser.add_argument('target_port', type=int, help='Target port')
    args = parser.parse_args()

    proxy = Proxy(args.bind_host, args.bind_port, args.target_host, args.target_port)

    asyncio.run(proxy.run())

if __name__ == '__main__':
    main()
