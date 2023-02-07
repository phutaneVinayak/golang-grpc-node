const path = require('path')
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const proto = protoLoader.loadSync(path.join(__dirname, '..', 'post_service.proto'));

const defination = grpc.loadPackageDefinition(proto);

const postList = [
    {id:1, title: "vinayak", text: "text 1"}
]

const getPosts = (call, cb) => {
    cb(null, {posts: postList})
}

const postPosts = (call, cb) => {
    setTimeout(() => {
        console.log("call", call.request);
        let userData =  call.request;
        userData["id"] = userData.id + 1;
        cb(null, userData)
    }, 180000);
}

const server = new grpc.Server();
server.addService(defination.PostService.service, {getPosts, postPosts});

server.bindAsync('localhost:8081', grpc.ServerCredentials.createInsecure(), port =>{
    console.log("GRPC server is started", port);
    server.start()
})