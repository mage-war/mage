package magpie

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kwunyeung/desmos/x/magpie/types"
	"github.com/rs/xid"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreatePost:
			return handleMsgCreatePost(ctx, keeper, msg)
		case MsgEditPost:
			return handleMsgEditPost(ctx, keeper, msg)
		case MsgLike:
			return handleMsgLike(ctx, keeper, msg)
		// case MsgUnlike:
		case MsgCreateSession:
			return handleMsgCreateSession(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Magpie Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle creating a new post
func handleMsgCreatePost(ctx sdk.Context, keeper Keeper, msg MsgCreatePost) sdk.Result {
	// if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
	// 	return sdk.ErrUnauthorized("Incorrect Owner").Result()
	// }
	post := Post{
		ID:            xid.New().String(),
		Message:       msg.Message,
		Created:       msg.Created,
		Likes:         0,
		Owner:         msg.Owner,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	err, success := keeper.SetPost(ctx, post)

	if err != nil {
		return err.Result()
	}

	if success {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCreatePost,
				sdk.NewAttribute(types.AttributeKeyPostID, post.ID),
				sdk.NewAttribute(types.AttributeKeyNamespace, post.Namespace),
				sdk.NewAttribute(types.AttributeKeyExternalOwner, post.ExternalOwner),
			),
		)
	}

	return sdk.Result{
		Data:   keeper.cdc.MustMarshalBinaryLengthPrefixed(post.ID),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgEditPost(ctx sdk.Context, keeper Keeper, msg MsgEditPost) sdk.Result {
	if !msg.Owner.Equals(keeper.GetPostOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	err, success := keeper.EditPost(ctx, msg.ID, msg.Message)

	if err != nil {
		return err.Result()
	}

	if success {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeEditPost,
				sdk.NewAttribute(types.AttributeKeyPostID, msg.ID),
			),
		)
	}

	return sdk.Result{
		Data:   keeper.cdc.MustMarshalBinaryLengthPrefixed(msg.ID),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgLike(ctx sdk.Context, keeper Keeper, msg MsgLike) sdk.Result {

	post := keeper.GetPost(ctx, msg.PostID)

	if msg.PostID != post.ID {
		return sdk.ErrUnknownRequest("Post doesn't exist").Result()
	}

	like := Like{
		ID:            xid.New().String(),
		Created:       msg.Created,
		PostID:        msg.PostID,
		Owner:         msg.Liker,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Liker.String()),
		),
	)

	err, success := keeper.SetLike(ctx, like.ID, like)

	if err != nil {
		return err.Result()
	}

	if success {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeLikePost,
				sdk.NewAttribute(types.AttributeKeyLikeID, like.ID),
				sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID),
				sdk.NewAttribute(types.AttributeKeyNamespace, msg.Namespace),
				sdk.NewAttribute(types.AttributeKeyExternalOwner, msg.ExternalOwner),
			),
		)
	}

	return sdk.Result{
		Data:   keeper.cdc.MustMarshalBinaryLengthPrefixed(like.ID),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgCreateSession(ctx sdk.Context, keeper Keeper, msg MsgCreateSession) sdk.Result {

	// query if a previous TX with the same namespace and external owner exists
	// if a query exists,
	// see if current time is between creation time and expiry time
	// if yes, then continue and emit event
	// else return error

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	// check if the signature is signed by the external address
	// addr, err := utils.GetAccAddressFromExternal(msg.ExternalOwner, msg.Namespace)

	// if err != nil {
	// 	return err.Result()
	// }

	// acc := auth.NewBaseAccountWithAddress(addr)

	// pubkey := acc.GetPubKey()

	// message := fmt.Sprintf(`{"account_number":"0","chain_id":"%s","fee":"","memo":"","msgs":[{"type":"desmos/MsgCreateSession","value":{"created":"%s","external_owner":"%s","namespace":"%s","owner":"%s","signature":null}}],"sequence":"0"}
	// `, ctx.ChainID(), msg.Created, msg.ExternalOwner, msg.Namespace, msg.Owner.String())

	// if !pubkey.VerifyBytes([]byte(message), []byte(msg.Signature)) {
	// 	return sdk.ErrUnauthorized("The session signature is not correct.").Result()
	// }

	session := Session{
		ID:            xid.New().String(),
		Created:       msg.Created,
		Expiry:        msg.Created.Add(time.Minute * 10),
		Owner:         msg.Owner,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
		Signature:     msg.Signature,
	}

	err, success := keeper.SetSession(ctx, session)

	if err != nil {
		return err.Result()
	}

	if success {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCreateSession,
				sdk.NewAttribute(types.AttributeKeySessionID, session.ID),
				sdk.NewAttribute(types.AttributeKeyNamespace, msg.Namespace),
				sdk.NewAttribute(types.AttributeKeyExternalOwner, msg.ExternalOwner),
			),
		)
	}

	return sdk.Result{
		Events: ctx.EventManager().Events(),
	}
}