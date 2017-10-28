# umbled
This is an IRC bot that sits in a channel and keeps the connection open. It
is different in that it is more accepting of errors over the lifetime of a
connection.

If it encounters an error, it notifies the channel about it the next time
it connects.

This is to try to see if such behaviour helps recovery. Specifically I want
to see if it is possible for retrying to make any difference in keeping a
connection alive in the face of an unreliable connection.

It is also to track issues.
