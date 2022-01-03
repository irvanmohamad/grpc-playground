"""The Python implementation of the GRPC helloworld.Greeter client."""

from __future__ import print_function

import logging
import time

import grpc
import greetpb.greet_pb2 as greet_pb2
import greetpb.greet_pb2_grpc as greet_grpc

_LOGGER = logging.getLogger(__name__)


# request/reply
def doUnary(stub=None):
    req = greet_pb2.GreetRequest(greeting={"first_name":"Sadhan", "last_name":"Sarker"})
    response = stub.Greet(req)
    print("Greeter client received: " + response.result)
    _LOGGER.error('Quota failure: %s', info)


# server sending streaming to client
def doServerStreaming(stub=None):
    req = greet_pb2.GreetRequest(greeting={"first_name":"Sadhan", "last_name":"Sarker"})
    response = stub.GreetManyTimesMethod(req)
    try:
       for r in response:
           print(f"Server sending streaming to client = {r.result}")
    except Exception as err:
        print(err)


# client sending stream to server
def doClientStreaming(stub=None):
    print("Starting to do a Client Streaming RPC...")
    response = stub.LongGreet(send_req())
    print(f"LongGreet Response:: {response.result}")


# client sending stream to server
def doClientStreaming_alt(stub=None):
    print("Starting to do a Client Streaming RPC...")

    def send_data():
        bulk_req = [
            {"first_name":"Sadhan", "last_name":"Sarker"},
            {"first_name":"Ripon", "last_name":"Sarker"},
        ]
        for req in bulk_req:
            print(f'Sending req: {req}', end="\n")
            yield greet_pb2.GreetEveryoneRequest(greeting=req)
            time.sleep(0.1)

    req_it = send_data()
    response = stub.LongGreet(req_it)
    print(f"LongGreet Response:: {response.result}")


def send_req():

    bulk_req = [
        {"first_name":"Sadhan", "last_name":"Sarker"},
        {"first_name":"Ripon", "last_name":"Sarker"},
        {"first_name":"Hannan", "last_name":"Sarker"},
        {"first_name":"Kharim", "last_name":"Sarker"}
    ]

    for req in bulk_req:
        print(f'Sending req: {req}', end="\n")
        yield greet_pb2.GreetEveryoneRequest(greeting=req)
        time.sleep(0.1)

    # while True:
    #     i = input("Enter a anything or 'q' to quit: ")
    #     if i == "q":
    #         break
    #     try:
    #         req = {"first_name":"Sadhan", "last_name":"Sarker"}
    #     except ValueError:
    #         continue
    #     yield greet_pb2.GreetEveryoneRequest(greeting=req)
    #     time.sleep(0.1)
    #     print(f'Sending req: {req}', end="\n")


# Bi-directional
def doStreamBoth(stub=None):
    print("Starting to do a Bi-directional Streaming RPC...")
    response = stub.GreetEveryoneMethod(send_req())
    try:
        for r in response:
            print(f"Response from GreetEveryone: {r.result}", end="\n")
    except Exception as err:
        print("Run server first")



def run():
    with grpc.insecure_channel('localhost:8080') as channel:
        stub = greet_grpc.GreetServiceStub(channel)

        #doUnary(stub)
        #doServerStreaming(stub)
        doClientStreaming(stub)
        #doStreamBoth(stub)



if __name__ == '__main__':
    logging.basicConfig()

    try:
        run()
    except Exception as e:
        print("Cant connect with server")