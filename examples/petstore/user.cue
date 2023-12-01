package petstore

// User is the user resource definition
resources: "User": {
	plural: "users"
	title: {
		zh: "用户"
		en: "User"
	}
	description: {
		zh: "用户信息"
		en: "User info"
	}
}

resources: "User": models: "User": {
	title: {
		zh: "用户"
		en: "User"
	}
	properties: [
		{
			name: "id"
			title: {
				zh: "ID"
				en: "ID"
			}
			schema: {
				type:     "integer"
				required: true
			}
		},
		{
			name: "info"
			title: {
				zh: "用户信息"
				en: "User info"
			}
			schema: {
				type:  "object"
				model: "UserInfo"
			}
		},
		{
			name: "friendship"
			title: {
				zh: "好友关系"
				en: "Friendship"
			}
			schema: {
				type: "array"
				elem: {
					type: "ref"
					ref:  "User"
				}
			}
		},
	]
}

resources: "User": models: "UserInfo": {
	title: {
		zh: "用户信息"
		en: "User info"
	}
	properties: [
		{
			name: "name"
			title: {
				zh: "名称"
				en: "Name"
			}
			schema: {
				type:     "string"
				required: true
			}
		},
		{
			name: "age"
			title: {
				zh: "年龄"
				en: "Age"
			}
			schema: {
				type:     "integer"
				required: true
			}
		},
	]
}
