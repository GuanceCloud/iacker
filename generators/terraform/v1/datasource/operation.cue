package datasource

import (
	gotemplate "text/template"

	template "github.com/GuanceCloud/iacker/pkg/template/v1"
	naming "github.com/GuanceCloud/iacker/pkg/helpers/naming"
)

_data_source_template: """
	// Code generated by Guance Cloud Code Generation Pipeline. DO NOT EDIT.
	
	package {{ .name.lower }}
	
	import (
		"context"
	    _ "embed"
	
		"github.com/hashicorp/go-multierror"
		"github.com/hashicorp/terraform-plugin-framework/datasource"
		"github.com/hashicorp/terraform-plugin-framework/types"
	
		"github.com/GuanceCloud/terraform-provider-guance/internal/helpers/tfcodec"
		"github.com/GuanceCloud/terraform-provider-guance/internal/consts"
		"github.com/GuanceCloud/terraform-provider-guance/internal/sdk"
	)
	
	//go:embed README.md
	var resourceDocument string
	
	// Ensure the implementation satisfies the expected interfaces.
	var (
		_ datasource.DataSource              = &{{ .name.lowercamel }}DataSource{}
		_ datasource.DataSourceWithConfigure = &{{ .name.lowercamel }}DataSource{}
	)
	
	// New{{ .name.uppercamel }}DataSource is a helper function to simplify the provider implementation.
	func New{{ .name.uppercamel }}DataSource() datasource.DataSource {
		return &{{ .name.lowercamel }}DataSource{}
	}
	
	// {{ .name.lowercamel }}DataSource is the data source implementation.
	type {{ .name.lowercamel }}DataSource struct {
		client *sdk.Client[sdk.Resource]
	}
	
	// Metadata returns the data source type name.
	func (d *{{ .name.lowercamel }}DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
		resp.TypeName = req.ProviderTypeName + "_{{ .name.lower }}"
	}
	
	// Schema defines the schema for the data source.
	func (d *{{ .name.lowercamel }}DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
		resp.Schema = dataSourceSchema
	}
	
	// Configure adds the provider configured client to the data source.
	func (d *{{ .name.lowercamel }}DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
		if req.ProviderData == nil {
			return
		}
	
		d.client = req.ProviderData.(*sdk.Client[sdk.Resource])
	}
	
	// Read refreshes the Terraform state with the latest data.
	func (d *{{ .name.lowercamel }}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
		var state {{ .rsname.lowercamel }}DataSourceModel
		diags := req.Config.Get(ctx, &state)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	
		results, err := d.client.List(ctx, &sdk.ListOptions{
			MaxResults: state.MaxResults.ValueInt64(),
			TypeName:   consts.TypeName{{ .rsname.uppercamel }},
		})
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to List",
				err.Error(),
			)
			return
		}
	
		var mErr error
		var items []*{{ .rsname.lowercamel }}ResourceModel
		for _, rd := range results.ResourceDescriptions {
			if !sdk.FilterAllSuccess(rd.Properties, state.Filters...) {
				continue
			}
	
			item := &{{ .rsname.lowercamel }}ResourceModel{}
			if err := tfcodec.DecodeJSON([]byte(rd.Properties), item); err != nil {
				mErr = multierror.Append(mErr, fmt.Errorf("unable to decode properties: %w", err))
				continue
			}
			item.SetId(rd.Identifier)
			item.SetCreatedAt(rd.CreatedAt)
	
			items = append(items, item)
		}
		if mErr != nil {
			resp.Diagnostics.AddError(
				"Unable to List resources",
				mErr.Error(),
			)
			return
		}
		state.Items = items
		state.ID = types.StringValue("placeholder")
	
		// Set state
		diags = resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	"""

for rsname, rsinfo in inputs.resources {
	if (*rsinfo.meta.datasource | false) {
		outputs: files: "internal/datasources/\((naming.#Snake & {name: rsinfo.plural}).lower)/datasource.go": template.#File & {
			content: gotemplate.Execute(_data_source_template, {
				"name":   naming.#Snake & {"name":      rsinfo.plural}
				"rsname": naming.#UpperCamel & {"name": rsname}
			})
		}
	}
}
