package main

import future.keywords.contains
import future.keywords.if
import future.keywords.in

# By default, deny requests.
default allow := false

# Allow admins to do anything.
# allow if user_is_admin

# Allow the action if the user is granted permission to perform the action.
allow if {
	# Find grants for the user.
	some grant in role_grants[input.role][input.resource_type]

	# Check if the grant permits the action.
	input.action == grant
}

# user_is_admin is true if "admin" is among the user's roles as per data.role_grants
# user_is_admin if input.role == "admin"


# Data
role_grants := {
	"admin": {
		"user": [
			"read",
			"write",
		],
		"private": [
			"read",
			"write",
			"delete",
		],
		"protected": [],
		"public": [],
	},
	"editor": {
		"user": [
			"read",
			"write",
		],
		"private": [
			"read",
			"write",
			"delete",
		],
		"protected": [],
		"public": [],
	},
	"viewer": {
		"user": [],
		"private": ["read"],
		"protected": [
			"read",
			"delete",
		],
		"public": [
			"read",
			"delete",
		],
	},
	"guest": {
		"user": [],
		"private": [],
		"protected": ["read"],
		"public": ["read"],
	},
}
