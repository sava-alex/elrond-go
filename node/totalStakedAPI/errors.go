package totalStakedAPI

import "errors"

// ErrInvalidTotalStakedValueCacheDuration signals that an invalid cache duration has been provided
var ErrInvalidTotalStakedValueCacheDuration = errors.New("invalid total staked value cache duration")

// ErrCannotCastAccountHandlerToUserAccount signal that returned account is wrong
var ErrCannotCastAccountHandlerToUserAccount = errors.New("cannot cast AccountHandler to UserAccount")

// ErrCannotReturnTotalStakedFromShardNode signals that total staked cannot be returned by a shard node
var ErrCannotReturnTotalStakedFromShardNode = errors.New("total staked value cannot be returned by a shard node")

// ErrNilMarshalizer signals that a nil marshalizer has been provided
var ErrNilMarshalizer = errors.New("trying to set nil marshalizer")

// ErrNilAccountsAdapter signals that a nil accounts adapter has been provided
var ErrNilAccountsAdapter = errors.New("trying to set nil accounts adapter")
