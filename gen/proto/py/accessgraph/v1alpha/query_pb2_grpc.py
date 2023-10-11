# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from accessgraph.v1alpha import query_pb2 as accessgraph_dot_v1alpha_dot_query__pb2


class AccessGraphServiceStub(object):
    """AccessGraphService is a service for interacting the access graph service.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Query = channel.unary_unary(
                '/accessgraph.v1alpha.AccessGraphService/Query',
                request_serializer=accessgraph_dot_v1alpha_dot_query__pb2.QueryRequest.SerializeToString,
                response_deserializer=accessgraph_dot_v1alpha_dot_query__pb2.QueryResponse.FromString,
                )
        self.GetFile = channel.unary_unary(
                '/accessgraph.v1alpha.AccessGraphService/GetFile',
                request_serializer=accessgraph_dot_v1alpha_dot_query__pb2.GetFileRequest.SerializeToString,
                response_deserializer=accessgraph_dot_v1alpha_dot_query__pb2.GetFileResponse.FromString,
                )


class AccessGraphServiceServicer(object):
    """AccessGraphService is a service for interacting the access graph service.
    """

    def Query(self, request, context):
        """Query queries the access graph.
        Currently only used by WebUI.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetFile(self, request, context):
        """GetFile gets a static UI file from the access graph container.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AccessGraphServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Query': grpc.unary_unary_rpc_method_handler(
                    servicer.Query,
                    request_deserializer=accessgraph_dot_v1alpha_dot_query__pb2.QueryRequest.FromString,
                    response_serializer=accessgraph_dot_v1alpha_dot_query__pb2.QueryResponse.SerializeToString,
            ),
            'GetFile': grpc.unary_unary_rpc_method_handler(
                    servicer.GetFile,
                    request_deserializer=accessgraph_dot_v1alpha_dot_query__pb2.GetFileRequest.FromString,
                    response_serializer=accessgraph_dot_v1alpha_dot_query__pb2.GetFileResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'accessgraph.v1alpha.AccessGraphService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AccessGraphService(object):
    """AccessGraphService is a service for interacting the access graph service.
    """

    @staticmethod
    def Query(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/accessgraph.v1alpha.AccessGraphService/Query',
            accessgraph_dot_v1alpha_dot_query__pb2.QueryRequest.SerializeToString,
            accessgraph_dot_v1alpha_dot_query__pb2.QueryResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/accessgraph.v1alpha.AccessGraphService/GetFile',
            accessgraph_dot_v1alpha_dot_query__pb2.GetFileRequest.SerializeToString,
            accessgraph_dot_v1alpha_dot_query__pb2.GetFileResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
