const messages = require('../services/greetpb/greet_pb')
const services = require('../services/greetpb/greet_grpc_pb')
const grpc = require('@grpc/grpc-js');
const _ = require('lodash');

const doUnary = (client) => {
    let greeting = new messages.Greeting();
    greeting.setFirstName('Sadhan');
    greeting.setLastName('Sarker');

    // console.log(greeting);

    let request = new messages.GreetRequest();
    request.setGreeting(greeting);

    client.greet(request, function (err, response) {
        if (response) console.log('Greeting: From Server', response.getResult());

        else console.log('Please run server first:', err);
    });
}


const doServerStreaming = (client) => {
    console.log('Client calling Server Streaming');

    let greeting1 = new messages.Greeting();
    greeting1.setFirstName('Sadhan');
    greeting1.setLastName('Sarker');

    let request = new messages.GreetManyTimesRequest();
    request.setGreeting(greeting1);

    // call method
    let call = client.greetManyTimesMethod(request);
    call.on('data', function (response) {
        console.log('Server Streaming Response:', response.getResult());
    });
    call.on('end', function () {
        console.log('All streaming have been end');
    });
}


const doClientStreaming = (client) => {
    console.log('Client Streaming to Server');

    let greeting1 = new messages.Greeting();
    greeting1.setFirstName('Sadhan');
    greeting1.setLastName('Sarker');

    let greeting2 = new messages.Greeting();
    greeting2.setFirstName('Sourav');
    greeting2.setLastName('Sarker');

    let request1 = new messages.LongGreetRequest();
    request1.setGreeting(greeting1);

    let request2 = new messages.LongGreetRequest();
    request2.setGreeting(greeting2);

    let requests = [request1, request2];

    let call = client.longGreet(function (error, response) {
        if (response) console.log('Server Response Payload:', response.getResult());
        else console.log('Please run server first:', error);
    });


    // both are same
    requests.forEach(function (request) {
        console.log('Client Sending Payload:', request.getGreeting().getFirstName());
        call.write(request);
    });
    // _.each(requests, function (response) { call.write(response);})

    call.end();
}


const doBidirectionalStreaming = (client) => {
    let call = client.greetEveryoneMethod();


    let greeting1 = new messages.Greeting();
    greeting1.setFirstName('Sadhan');
    greeting1.setLastName('Sarker');

    let greeting2 = new messages.Greeting();
    greeting2.setFirstName('Sourav');
    greeting2.setLastName('Sarker');

    let request1 = new messages.GreetEveryoneRequest();
    request1.setGreeting(greeting1);

    let request2 = new messages.GreetEveryoneRequest();
    request2.setGreeting(greeting2);

    let requests = [request1, request2];

    requests.forEach(function (request) {
        console.log('Client Sending Payload:', request.getGreeting().getFirstName());
        call.write(request);    // send request to server
    });

    // read server side streaming response
    call.on('data', function (response) {
        console.log('Server Response Payload:', response.getResult());
    });

    call.on('end', function () {
        console.log('All streaming have been end');
    });
    call.end();
}

function main() {

    let client = new services.GreetServiceClient('localhost:8080', grpc.credentials.createInsecure());
    doUnary(client);
    // doServerStreaming(client);
    // doClientStreaming(client);
    // doBidirectionalStreaming(client);
}

// https://stackoverflow.com/questions/62483102/how-to-stream-bytes-using-grpc-in-nodejs
// https://github.com/aditya-sridhar/grpc-streams-nodejs-demo

main();