/**
 * @fileoverview gRPC-Web generated client stub for {{ .ProjectName }}
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.{{ .ProjectName }} = require('./{{ .ProjectName }}_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.{{ .ProjectName }}.TodoRpcClient =
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
proto.{{ .ProjectName }}.TodoRpcPromiseClient =
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
 *   !proto.{{ .ProjectName }}.SearchRequest,
 *   !proto.{{ .ProjectName }}.SearchResponse>}
 */
const methodDescriptor_TodoRpc_Search = new grpc.web.MethodDescriptor(
  '/{{ .ProjectName }}.TodoRpc/Search',
  grpc.web.MethodType.UNARY,
  proto.{{ .ProjectName }}.SearchRequest,
  proto.{{ .ProjectName }}.SearchResponse,
  /**
   * @param {!proto.{{ .ProjectName }}.SearchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.{{ .ProjectName }}.SearchResponse.deserializeBinary
);


/**
 * @param {!proto.{{ .ProjectName }}.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.{{ .ProjectName }}.SearchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.{{ .ProjectName }}.SearchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.{{ .ProjectName }}.TodoRpcClient.prototype.search =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/{{ .ProjectName }}.TodoRpc/Search',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Search,
      callback);
};


/**
 * @param {!proto.{{ .ProjectName }}.SearchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.{{ .ProjectName }}.SearchResponse>}
 *     Promise that resolves to the response
 */
proto.{{ .ProjectName }}.TodoRpcPromiseClient.prototype.search =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/{{ .ProjectName }}.TodoRpc/Search',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Search);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.{{ .ProjectName }}.TodoItem,
 *   !proto.{{ .ProjectName }}.TodoItem>}
 */
const methodDescriptor_TodoRpc_Create = new grpc.web.MethodDescriptor(
  '/{{ .ProjectName }}.TodoRpc/Create',
  grpc.web.MethodType.UNARY,
  proto.{{ .ProjectName }}.TodoItem,
  proto.{{ .ProjectName }}.TodoItem,
  /**
   * @param {!proto.{{ .ProjectName }}.TodoItem} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.{{ .ProjectName }}.TodoItem.deserializeBinary
);


/**
 * @param {!proto.{{ .ProjectName }}.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.{{ .ProjectName }}.TodoItem)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.{{ .ProjectName }}.TodoItem>|undefined}
 *     The XHR Node Readable Stream
 */
proto.{{ .ProjectName }}.TodoRpcClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/{{ .ProjectName }}.TodoRpc/Create',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Create,
      callback);
};


/**
 * @param {!proto.{{ .ProjectName }}.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.{{ .ProjectName }}.TodoItem>}
 *     Promise that resolves to the response
 */
proto.{{ .ProjectName }}.TodoRpcPromiseClient.prototype.create =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/{{ .ProjectName }}.TodoRpc/Create',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Create);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.{{ .ProjectName }}.TodoItem,
 *   !proto.{{ .ProjectName }}.TodoItem>}
 */
const methodDescriptor_TodoRpc_Update = new grpc.web.MethodDescriptor(
  '/{{ .ProjectName }}.TodoRpc/Update',
  grpc.web.MethodType.UNARY,
  proto.{{ .ProjectName }}.TodoItem,
  proto.{{ .ProjectName }}.TodoItem,
  /**
   * @param {!proto.{{ .ProjectName }}.TodoItem} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.{{ .ProjectName }}.TodoItem.deserializeBinary
);


/**
 * @param {!proto.{{ .ProjectName }}.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.{{ .ProjectName }}.TodoItem)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.{{ .ProjectName }}.TodoItem>|undefined}
 *     The XHR Node Readable Stream
 */
proto.{{ .ProjectName }}.TodoRpcClient.prototype.update =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/{{ .ProjectName }}.TodoRpc/Update',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Update,
      callback);
};


/**
 * @param {!proto.{{ .ProjectName }}.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.{{ .ProjectName }}.TodoItem>}
 *     Promise that resolves to the response
 */
proto.{{ .ProjectName }}.TodoRpcPromiseClient.prototype.update =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/{{ .ProjectName }}.TodoRpc/Update',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Update);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.{{ .ProjectName }}.TodoItem,
 *   !proto.{{ .ProjectName }}.Nothing>}
 */
const methodDescriptor_TodoRpc_Remove = new grpc.web.MethodDescriptor(
  '/{{ .ProjectName }}.TodoRpc/Remove',
  grpc.web.MethodType.UNARY,
  proto.{{ .ProjectName }}.TodoItem,
  proto.{{ .ProjectName }}.Nothing,
  /**
   * @param {!proto.{{ .ProjectName }}.TodoItem} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.{{ .ProjectName }}.Nothing.deserializeBinary
);


/**
 * @param {!proto.{{ .ProjectName }}.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.{{ .ProjectName }}.Nothing)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.{{ .ProjectName }}.Nothing>|undefined}
 *     The XHR Node Readable Stream
 */
proto.{{ .ProjectName }}.TodoRpcClient.prototype.remove =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/{{ .ProjectName }}.TodoRpc/Remove',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Remove,
      callback);
};


/**
 * @param {!proto.{{ .ProjectName }}.TodoItem} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.{{ .ProjectName }}.Nothing>}
 *     Promise that resolves to the response
 */
proto.{{ .ProjectName }}.TodoRpcPromiseClient.prototype.remove =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/{{ .ProjectName }}.TodoRpc/Remove',
      request,
      metadata || {},
      methodDescriptor_TodoRpc_Remove);
};


module.exports = proto.{{ .ProjectName }};

