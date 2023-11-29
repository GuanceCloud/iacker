package petstore

// Pet is the pet resource definition
resources: "Pet": {
	plural: "pets"
	title: {
		zh: "宠物"
		en: "Pet"
	}
	description: {
		zh: "宠物很可爱，包括猫猫狗狗等"
		en: "Pets are cute, including cats, dogs, etc."
	}
}

// Pet model
resources: "Pet": models: "Pet": {
	title: {
		zh: "宠物"
		en: "Pet"
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
			name: "tags"
			title: {
				zh: "标签"
				en: "Tag"
			}
			schema: {
				type: "array"
				elem: {
					type: "string"
				}
			}
		},
		{
			name: "owner"
			title: {
				zh: "拥有者"
				en: "Owner"
			}
			schema: {
				type: "ref"
				ref:  "User"
			}
		},
		{
			name: "status"
			title: {
				zh: "状态"
				en: "Status"
			}
			schema: {
				type: "string"
				enum: [
					{
						name:  "AVAILABLE"
						value: "available"
						title: {
							zh: "可用"
							en: "Available"
						}
					},
					{
						name:  "PENDING"
						value: "pending"
						title: {
							zh: "待定"
							en: "Pending"
						}
					},
					{
						name:  "SOLD"
						value: "sold"
						title: {
							zh: "已售"
							en: "Sold"
						}
					},
				]
			}
		},
		{
			name: "extras"
			title: {
				zh: "额外信息"
				en: "Extra info"
			}
			schema: {
				type: "array"
				elem: {
					type:  "object"
					model: "ExtraInfo"
				}
			}
		},
	]
}

resources: "Pet": models: "ExtraInfo": {
	title: {
		zh: "额外信息"
		en: "Extra info"
	}
	properties: [
		{
			name: "key"
			title: {
				zh: "键"
				en: "Key"
			}
			schema: {
				type:     "string"
				required: true
			}
		},
		{
			name: "value"
			title: {
				zh: "值"
				en: "Value"
			}
			schema: {
				type:     "string"
				required: true
			}
		},
	]
}
