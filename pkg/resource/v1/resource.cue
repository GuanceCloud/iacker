package v1

#Resource: {
	plural?:      string @protobuf(1,string)
	title?:       #I18n  @protobuf(2,I18n)
	description?: #I18n  @protobuf(3,I18n)
	models?: {
		[string]: #Model
	} @protobuf(4,map[string]Model)
	meta?:       #Meta       @protobuf(5,Meta)
	identifier?: #Identifier @protobuf(6,Identifier)
	errors?: [...string] @protobuf(7,string)
}

#Identifier: {
	primary?: [...string] @protobuf(1,string)
}

#Meta: {
	datasource?: bool @protobuf(1,bool)
}

#Model: {
	title?:       #I18n @protobuf(1,I18n)
	description?: #I18n @protobuf(2,I18n)
	properties?: [...#Property] @protobuf(3,Property)
}

#Property: {
	name?:        string      @protobuf(1,string)
	title?:       #I18n       @protobuf(2,I18n)
	description?: #I18n       @protobuf(3,I18n)
	schema?:      #PropSchema @protobuf(4,PropSchema)
	meta?:        #PropMeta   @protobuf(5,PropMeta)
}

#I18n: {
	zh?: string @protobuf(1,string)
	en?: string @protobuf(2,string)
}

#PropSchema: {
	type?:     string      @protobuf(1,string)
	format?:   string      @protobuf(2,string)
	required?: bool        @protobuf(3,bool)
	elem?:     #ElemSchema @protobuf(4,ElemSchema)
	enum?: [...#Enum] @protobuf(5,Enum)
	model?: string @protobuf(6,string)
	ref?:   string @protobuf(7,string)
}

#ElemSchema: {
	type?:   string @protobuf(1,string)
	format?: string @protobuf(2,string)
	model?:  string @protobuf(3,string)
	ref?:    string @protobuf(4,string)
	enum?: [...#Enum] @protobuf(5,Enum)
}

#PropMeta: {
	dynamic?:   bool @protobuf(1,bool)
	immutable?: bool @protobuf(2,bool)
}

#Enum: {
	name?:  string @protobuf(1,string)
	value?: string @protobuf(2,string)
	title?: #I18n  @protobuf(3,I18n)
}

#Error: {
	title?: #I18n      @protobuf(1,I18n)
	meta?:  #ErrorMeta @protobuf(2,ErrorMeta)
}

#ErrorMeta: {
	code?: int64 @protobuf(1,int64)
}
