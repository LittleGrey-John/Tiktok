// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	douyincomment "douyin/kitex_gen/douyincomment"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CommentAction(ctx context.Context, request *douyincomment.DouyinCommentActionRequest, uid int64, callOptions ...callopt.Option) (r *douyincomment.DouyinCommentActionResponse, err error)
	CommentList(ctx context.Context, request *douyincomment.DouyinCommentListRequest, callOptions ...callopt.Option) (r *douyincomment.DouyinCommentListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) CommentAction(ctx context.Context, request *douyincomment.DouyinCommentActionRequest, uid int64, callOptions ...callopt.Option) (r *douyincomment.DouyinCommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, request, uid)
}

func (p *kCommentServiceClient) CommentList(ctx context.Context, request *douyincomment.DouyinCommentListRequest, callOptions ...callopt.Option) (r *douyincomment.DouyinCommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, request)
}
