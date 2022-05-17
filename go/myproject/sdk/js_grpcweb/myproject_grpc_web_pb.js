/**
 * @fileoverview gRPC-Web generated client stub for myproject
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.myproject = require('./myproject_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.myproject.TodoRpcClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.myproject.TodoRpcPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.myproject.SearchRequest,
 *   !proto.myproject.SearchResponse>}
 */
const methodDescriptor_TodoRpc_Search = new grpc.web.MethodDescriptor(
  '/myproject.TodoRpc/Search',
  grpc.web.MethodType.UNARY,
  proto.myproject.SearchRequest,
  proto.myproject.SearchResponse,
  /**
   * @param {!proto.myproject.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.myproject.SearchResponse.deserializeBinary
);


/**
 * @param {!proto.myproject.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.myproject.SearchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.myproject.SearchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.myproject.TodoRpcClient.prototype.search =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/myproject.TodoRpc/Search',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Search,
      callback);
};


/**
 * @param {!proto.myproject.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.myproject.SearchResponse>}
 *     Promise that resolves to the response
 */
proto.myproject.TodoRpcPromiseClient.prototype.search =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/myproject.TodoRpc/Search',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Search);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.myproject.TodoItem,
 *   !proto.myproject.TodoItem>}
 */
const methodDescriptor_TodoRpc_Create = new grpc.web.MethodDescriptor(
  '/myproject.TodoRpc/Create',
  grpc.web.MethodType.UNARY,
  proto.myproject.TodoItem,
  proto.myproject.TodoItem,
  /**
   * @param {!proto.myproject.TodoItem} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.myproject.TodoItem.deserializeBinary
);


/**
 * @param {!proto.myproject.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.myproject.TodoItem)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.myproject.TodoItem>|undefined}
 *     The XHR Node Readable Stream
 */
proto.myproject.TodoRpcClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/myproject.TodoRpc/Create',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Create,
      callback);
};


/**
 * @param {!proto.myproject.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.myproject.TodoItem>}
 *     Promise that resolves to the response
 */
proto.myproject.TodoRpcPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/myproject.TodoRpc/Create',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.myproject.TodoItem,
 *   !proto.myproject.TodoItem>}
 */
const methodDescriptor_TodoRpc_Update = new grpc.web.MethodDescriptor(
  '/myproject.TodoRpc/Update',
  grpc.web.MethodType.UNARY,
  proto.myproject.TodoItem,
  proto.myproject.TodoItem,
  /**
   * @param {!proto.myproject.TodoItem} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.myproject.TodoItem.deserializeBinary
);


/**
 * @param {!proto.myproject.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.myproject.TodoItem)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.myproject.TodoItem>|undefined}
 *     The XHR Node Readable Stream
 */
proto.myproject.TodoRpcClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/myproject.TodoRpc/Update',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Update,
      callback);
};


/**
 * @param {!proto.myproject.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.myproject.TodoItem>}
 *     Promise that resolves to the response
 */
proto.myproject.TodoRpcPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/myproject.TodoRpc/Update',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.myproject.TodoItem,
 *   !proto.myproject.Nothing>}
 */
const methodDescriptor_TodoRpc_Remove = new grpc.web.MethodDescriptor(
  '/myproject.TodoRpc/Remove',
  grpc.web.MethodType.UNARY,
  proto.myproject.TodoItem,
  proto.myproject.Nothing,
  /**
   * @param {!proto.myproject.TodoItem} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.myproject.Nothing.deserializeBinary
);


/**
 * @param {!proto.myproject.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.myproject.Nothing)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.myproject.Nothing>|undefined}
 *     The XHR Node Readable Stream
 */
proto.myproject.TodoRpcClient.prototype.remove =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/myproject.TodoRpc/Remove',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Remove,
      callback);
};


/**
 * @param {!proto.myproject.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.myproject.Nothing>}
 *     Promise that resolves to the response
 */
proto.myproject.TodoRpcPromiseClient.prototype.remove =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/myproject.TodoRpc/Remove',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Remove);
};


module.exports = proto.myproject;

