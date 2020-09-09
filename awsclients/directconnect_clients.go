package awsclients

import (
	"github.com/aws/aws-sdk-go/service/directconnect"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type DirectConnectClient interface {
       AcceptDirectConnectGatewayAssociationProposal(ctx workflow.Context, input *directconnect.AcceptDirectConnectGatewayAssociationProposalInput) (*directconnect.AcceptDirectConnectGatewayAssociationProposalOutput, error)
       AcceptDirectConnectGatewayAssociationProposalAsync(ctx workflow.Context, input *directconnect.AcceptDirectConnectGatewayAssociationProposalInput) *DirectconnectAcceptDirectConnectGatewayAssociationProposalResult

       AllocateConnectionOnInterconnect(ctx workflow.Context, input *directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error)
       AllocateConnectionOnInterconnectAsync(ctx workflow.Context, input *directconnect.AllocateConnectionOnInterconnectInput) *DirectconnectAllocateConnectionOnInterconnectResult

       AllocateHostedConnection(ctx workflow.Context, input *directconnect.AllocateHostedConnectionInput) (*directconnect.Connection, error)
       AllocateHostedConnectionAsync(ctx workflow.Context, input *directconnect.AllocateHostedConnectionInput) *DirectconnectAllocateHostedConnectionResult

       AllocatePrivateVirtualInterface(ctx workflow.Context, input *directconnect.AllocatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
       AllocatePrivateVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.AllocatePrivateVirtualInterfaceInput) *DirectconnectAllocatePrivateVirtualInterfaceResult

       AllocatePublicVirtualInterface(ctx workflow.Context, input *directconnect.AllocatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
       AllocatePublicVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.AllocatePublicVirtualInterfaceInput) *DirectconnectAllocatePublicVirtualInterfaceResult

       AllocateTransitVirtualInterface(ctx workflow.Context, input *directconnect.AllocateTransitVirtualInterfaceInput) (*directconnect.AllocateTransitVirtualInterfaceOutput, error)
       AllocateTransitVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.AllocateTransitVirtualInterfaceInput) *DirectconnectAllocateTransitVirtualInterfaceResult

       AssociateConnectionWithLag(ctx workflow.Context, input *directconnect.AssociateConnectionWithLagInput) (*directconnect.Connection, error)
       AssociateConnectionWithLagAsync(ctx workflow.Context, input *directconnect.AssociateConnectionWithLagInput) *DirectconnectAssociateConnectionWithLagResult

       AssociateHostedConnection(ctx workflow.Context, input *directconnect.AssociateHostedConnectionInput) (*directconnect.Connection, error)
       AssociateHostedConnectionAsync(ctx workflow.Context, input *directconnect.AssociateHostedConnectionInput) *DirectconnectAssociateHostedConnectionResult

       AssociateVirtualInterface(ctx workflow.Context, input *directconnect.AssociateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
       AssociateVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.AssociateVirtualInterfaceInput) *DirectconnectAssociateVirtualInterfaceResult

       ConfirmConnection(ctx workflow.Context, input *directconnect.ConfirmConnectionInput) (*directconnect.ConfirmConnectionOutput, error)
       ConfirmConnectionAsync(ctx workflow.Context, input *directconnect.ConfirmConnectionInput) *DirectconnectConfirmConnectionResult

       ConfirmPrivateVirtualInterface(ctx workflow.Context, input *directconnect.ConfirmPrivateVirtualInterfaceInput) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error)
       ConfirmPrivateVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.ConfirmPrivateVirtualInterfaceInput) *DirectconnectConfirmPrivateVirtualInterfaceResult

       ConfirmPublicVirtualInterface(ctx workflow.Context, input *directconnect.ConfirmPublicVirtualInterfaceInput) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error)
       ConfirmPublicVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.ConfirmPublicVirtualInterfaceInput) *DirectconnectConfirmPublicVirtualInterfaceResult

       ConfirmTransitVirtualInterface(ctx workflow.Context, input *directconnect.ConfirmTransitVirtualInterfaceInput) (*directconnect.ConfirmTransitVirtualInterfaceOutput, error)
       ConfirmTransitVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.ConfirmTransitVirtualInterfaceInput) *DirectconnectConfirmTransitVirtualInterfaceResult

       CreateBGPPeer(ctx workflow.Context, input *directconnect.CreateBGPPeerInput) (*directconnect.CreateBGPPeerOutput, error)
       CreateBGPPeerAsync(ctx workflow.Context, input *directconnect.CreateBGPPeerInput) *DirectconnectCreateBGPPeerResult

       CreateConnection(ctx workflow.Context, input *directconnect.CreateConnectionInput) (*directconnect.Connection, error)
       CreateConnectionAsync(ctx workflow.Context, input *directconnect.CreateConnectionInput) *DirectconnectCreateConnectionResult

       CreateDirectConnectGateway(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayInput) (*directconnect.CreateDirectConnectGatewayOutput, error)
       CreateDirectConnectGatewayAsync(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayInput) *DirectconnectCreateDirectConnectGatewayResult

       CreateDirectConnectGatewayAssociation(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayAssociationInput) (*directconnect.CreateDirectConnectGatewayAssociationOutput, error)
       CreateDirectConnectGatewayAssociationAsync(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayAssociationInput) *DirectconnectCreateDirectConnectGatewayAssociationResult

       CreateDirectConnectGatewayAssociationProposal(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayAssociationProposalInput) (*directconnect.CreateDirectConnectGatewayAssociationProposalOutput, error)
       CreateDirectConnectGatewayAssociationProposalAsync(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayAssociationProposalInput) *DirectconnectCreateDirectConnectGatewayAssociationProposalResult

       CreateInterconnect(ctx workflow.Context, input *directconnect.CreateInterconnectInput) (*directconnect.Interconnect, error)
       CreateInterconnectAsync(ctx workflow.Context, input *directconnect.CreateInterconnectInput) *DirectconnectCreateInterconnectResult

       CreateLag(ctx workflow.Context, input *directconnect.CreateLagInput) (*directconnect.Lag, error)
       CreateLagAsync(ctx workflow.Context, input *directconnect.CreateLagInput) *DirectconnectCreateLagResult

       CreatePrivateVirtualInterface(ctx workflow.Context, input *directconnect.CreatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
       CreatePrivateVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.CreatePrivateVirtualInterfaceInput) *DirectconnectCreatePrivateVirtualInterfaceResult

       CreatePublicVirtualInterface(ctx workflow.Context, input *directconnect.CreatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error)
       CreatePublicVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.CreatePublicVirtualInterfaceInput) *DirectconnectCreatePublicVirtualInterfaceResult

       CreateTransitVirtualInterface(ctx workflow.Context, input *directconnect.CreateTransitVirtualInterfaceInput) (*directconnect.CreateTransitVirtualInterfaceOutput, error)
       CreateTransitVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.CreateTransitVirtualInterfaceInput) *DirectconnectCreateTransitVirtualInterfaceResult

       DeleteBGPPeer(ctx workflow.Context, input *directconnect.DeleteBGPPeerInput) (*directconnect.DeleteBGPPeerOutput, error)
       DeleteBGPPeerAsync(ctx workflow.Context, input *directconnect.DeleteBGPPeerInput) *DirectconnectDeleteBGPPeerResult

       DeleteConnection(ctx workflow.Context, input *directconnect.DeleteConnectionInput) (*directconnect.Connection, error)
       DeleteConnectionAsync(ctx workflow.Context, input *directconnect.DeleteConnectionInput) *DirectconnectDeleteConnectionResult

       DeleteDirectConnectGateway(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayInput) (*directconnect.DeleteDirectConnectGatewayOutput, error)
       DeleteDirectConnectGatewayAsync(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayInput) *DirectconnectDeleteDirectConnectGatewayResult

       DeleteDirectConnectGatewayAssociation(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayAssociationInput) (*directconnect.DeleteDirectConnectGatewayAssociationOutput, error)
       DeleteDirectConnectGatewayAssociationAsync(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayAssociationInput) *DirectconnectDeleteDirectConnectGatewayAssociationResult

       DeleteDirectConnectGatewayAssociationProposal(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayAssociationProposalInput) (*directconnect.DeleteDirectConnectGatewayAssociationProposalOutput, error)
       DeleteDirectConnectGatewayAssociationProposalAsync(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayAssociationProposalInput) *DirectconnectDeleteDirectConnectGatewayAssociationProposalResult

       DeleteInterconnect(ctx workflow.Context, input *directconnect.DeleteInterconnectInput) (*directconnect.DeleteInterconnectOutput, error)
       DeleteInterconnectAsync(ctx workflow.Context, input *directconnect.DeleteInterconnectInput) *DirectconnectDeleteInterconnectResult

       DeleteLag(ctx workflow.Context, input *directconnect.DeleteLagInput) (*directconnect.Lag, error)
       DeleteLagAsync(ctx workflow.Context, input *directconnect.DeleteLagInput) *DirectconnectDeleteLagResult

       DeleteVirtualInterface(ctx workflow.Context, input *directconnect.DeleteVirtualInterfaceInput) (*directconnect.DeleteVirtualInterfaceOutput, error)
       DeleteVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.DeleteVirtualInterfaceInput) *DirectconnectDeleteVirtualInterfaceResult

       DescribeConnectionLoa(ctx workflow.Context, input *directconnect.DescribeConnectionLoaInput) (*directconnect.DescribeConnectionLoaOutput, error)
       DescribeConnectionLoaAsync(ctx workflow.Context, input *directconnect.DescribeConnectionLoaInput) *DirectconnectDescribeConnectionLoaResult

       DescribeConnections(ctx workflow.Context, input *directconnect.DescribeConnectionsInput) (*directconnect.Connections, error)
       DescribeConnectionsAsync(ctx workflow.Context, input *directconnect.DescribeConnectionsInput) *DirectconnectDescribeConnectionsResult

       DescribeConnectionsOnInterconnect(ctx workflow.Context, input *directconnect.DescribeConnectionsOnInterconnectInput) (*directconnect.Connections, error)
       DescribeConnectionsOnInterconnectAsync(ctx workflow.Context, input *directconnect.DescribeConnectionsOnInterconnectInput) *DirectconnectDescribeConnectionsOnInterconnectResult

       DescribeDirectConnectGatewayAssociationProposals(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error)
       DescribeDirectConnectGatewayAssociationProposalsAsync(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) *DirectconnectDescribeDirectConnectGatewayAssociationProposalsResult

       DescribeDirectConnectGatewayAssociations(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAssociationsInput) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error)
       DescribeDirectConnectGatewayAssociationsAsync(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAssociationsInput) *DirectconnectDescribeDirectConnectGatewayAssociationsResult

       DescribeDirectConnectGatewayAttachments(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAttachmentsInput) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error)
       DescribeDirectConnectGatewayAttachmentsAsync(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAttachmentsInput) *DirectconnectDescribeDirectConnectGatewayAttachmentsResult

       DescribeDirectConnectGateways(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewaysInput) (*directconnect.DescribeDirectConnectGatewaysOutput, error)
       DescribeDirectConnectGatewaysAsync(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewaysInput) *DirectconnectDescribeDirectConnectGatewaysResult

       DescribeHostedConnections(ctx workflow.Context, input *directconnect.DescribeHostedConnectionsInput) (*directconnect.Connections, error)
       DescribeHostedConnectionsAsync(ctx workflow.Context, input *directconnect.DescribeHostedConnectionsInput) *DirectconnectDescribeHostedConnectionsResult

       DescribeInterconnectLoa(ctx workflow.Context, input *directconnect.DescribeInterconnectLoaInput) (*directconnect.DescribeInterconnectLoaOutput, error)
       DescribeInterconnectLoaAsync(ctx workflow.Context, input *directconnect.DescribeInterconnectLoaInput) *DirectconnectDescribeInterconnectLoaResult

       DescribeInterconnects(ctx workflow.Context, input *directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error)
       DescribeInterconnectsAsync(ctx workflow.Context, input *directconnect.DescribeInterconnectsInput) *DirectconnectDescribeInterconnectsResult

       DescribeLags(ctx workflow.Context, input *directconnect.DescribeLagsInput) (*directconnect.DescribeLagsOutput, error)
       DescribeLagsAsync(ctx workflow.Context, input *directconnect.DescribeLagsInput) *DirectconnectDescribeLagsResult

       DescribeLoa(ctx workflow.Context, input *directconnect.DescribeLoaInput) (*directconnect.Loa, error)
       DescribeLoaAsync(ctx workflow.Context, input *directconnect.DescribeLoaInput) *DirectconnectDescribeLoaResult

       DescribeLocations(ctx workflow.Context, input *directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error)
       DescribeLocationsAsync(ctx workflow.Context, input *directconnect.DescribeLocationsInput) *DirectconnectDescribeLocationsResult

       DescribeTags(ctx workflow.Context, input *directconnect.DescribeTagsInput) (*directconnect.DescribeTagsOutput, error)
       DescribeTagsAsync(ctx workflow.Context, input *directconnect.DescribeTagsInput) *DirectconnectDescribeTagsResult

       DescribeVirtualGateways(ctx workflow.Context, input *directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error)
       DescribeVirtualGatewaysAsync(ctx workflow.Context, input *directconnect.DescribeVirtualGatewaysInput) *DirectconnectDescribeVirtualGatewaysResult

       DescribeVirtualInterfaces(ctx workflow.Context, input *directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error)
       DescribeVirtualInterfacesAsync(ctx workflow.Context, input *directconnect.DescribeVirtualInterfacesInput) *DirectconnectDescribeVirtualInterfacesResult

       DisassociateConnectionFromLag(ctx workflow.Context, input *directconnect.DisassociateConnectionFromLagInput) (*directconnect.Connection, error)
       DisassociateConnectionFromLagAsync(ctx workflow.Context, input *directconnect.DisassociateConnectionFromLagInput) *DirectconnectDisassociateConnectionFromLagResult

       ListVirtualInterfaceTestHistory(ctx workflow.Context, input *directconnect.ListVirtualInterfaceTestHistoryInput) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error)
       ListVirtualInterfaceTestHistoryAsync(ctx workflow.Context, input *directconnect.ListVirtualInterfaceTestHistoryInput) *DirectconnectListVirtualInterfaceTestHistoryResult

       StartBgpFailoverTest(ctx workflow.Context, input *directconnect.StartBgpFailoverTestInput) (*directconnect.StartBgpFailoverTestOutput, error)
       StartBgpFailoverTestAsync(ctx workflow.Context, input *directconnect.StartBgpFailoverTestInput) *DirectconnectStartBgpFailoverTestResult

       StopBgpFailoverTest(ctx workflow.Context, input *directconnect.StopBgpFailoverTestInput) (*directconnect.StopBgpFailoverTestOutput, error)
       StopBgpFailoverTestAsync(ctx workflow.Context, input *directconnect.StopBgpFailoverTestInput) *DirectconnectStopBgpFailoverTestResult

       TagResource(ctx workflow.Context, input *directconnect.TagResourceInput) (*directconnect.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *directconnect.TagResourceInput) *DirectconnectTagResourceResult

       UntagResource(ctx workflow.Context, input *directconnect.UntagResourceInput) (*directconnect.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *directconnect.UntagResourceInput) *DirectconnectUntagResourceResult

       UpdateDirectConnectGatewayAssociation(ctx workflow.Context, input *directconnect.UpdateDirectConnectGatewayAssociationInput) (*directconnect.UpdateDirectConnectGatewayAssociationOutput, error)
       UpdateDirectConnectGatewayAssociationAsync(ctx workflow.Context, input *directconnect.UpdateDirectConnectGatewayAssociationInput) *DirectconnectUpdateDirectConnectGatewayAssociationResult

       UpdateLag(ctx workflow.Context, input *directconnect.UpdateLagInput) (*directconnect.Lag, error)
       UpdateLagAsync(ctx workflow.Context, input *directconnect.UpdateLagInput) *DirectconnectUpdateLagResult

       UpdateVirtualInterfaceAttributes(ctx workflow.Context, input *directconnect.UpdateVirtualInterfaceAttributesInput) (*directconnect.UpdateVirtualInterfaceAttributesOutput, error)
       UpdateVirtualInterfaceAttributesAsync(ctx workflow.Context, input *directconnect.UpdateVirtualInterfaceAttributesInput) *DirectconnectUpdateVirtualInterfaceAttributesResult
}

type DirectconnectAcceptDirectConnectGatewayAssociationProposalResult struct {
	Result workflow.Future
}

func (r *DirectconnectAcceptDirectConnectGatewayAssociationProposalResult) Get(ctx workflow.Context) (*directconnect.AcceptDirectConnectGatewayAssociationProposalOutput, error) {
    var output directconnect.AcceptDirectConnectGatewayAssociationProposalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectAllocateConnectionOnInterconnectResult struct {
	Result workflow.Future
}

func (r *DirectconnectAllocateConnectionOnInterconnectResult) Get(ctx workflow.Context) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectAllocateHostedConnectionResult struct {
	Result workflow.Future
}

func (r *DirectconnectAllocateHostedConnectionResult) Get(ctx workflow.Context) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectAllocatePrivateVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectAllocatePrivateVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectAllocatePublicVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectAllocatePublicVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectAllocateTransitVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectAllocateTransitVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.AllocateTransitVirtualInterfaceOutput, error) {
    var output directconnect.AllocateTransitVirtualInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectAssociateConnectionWithLagResult struct {
	Result workflow.Future
}

func (r *DirectconnectAssociateConnectionWithLagResult) Get(ctx workflow.Context) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectAssociateHostedConnectionResult struct {
	Result workflow.Future
}

func (r *DirectconnectAssociateHostedConnectionResult) Get(ctx workflow.Context) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectAssociateVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectAssociateVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectConfirmConnectionResult struct {
	Result workflow.Future
}

func (r *DirectconnectConfirmConnectionResult) Get(ctx workflow.Context) (*directconnect.ConfirmConnectionOutput, error) {
    var output directconnect.ConfirmConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectConfirmPrivateVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectConfirmPrivateVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error) {
    var output directconnect.ConfirmPrivateVirtualInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectConfirmPublicVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectConfirmPublicVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error) {
    var output directconnect.ConfirmPublicVirtualInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectConfirmTransitVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectConfirmTransitVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.ConfirmTransitVirtualInterfaceOutput, error) {
    var output directconnect.ConfirmTransitVirtualInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreateBGPPeerResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreateBGPPeerResult) Get(ctx workflow.Context) (*directconnect.CreateBGPPeerOutput, error) {
    var output directconnect.CreateBGPPeerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreateConnectionResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreateConnectionResult) Get(ctx workflow.Context) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreateDirectConnectGatewayResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreateDirectConnectGatewayResult) Get(ctx workflow.Context) (*directconnect.CreateDirectConnectGatewayOutput, error) {
    var output directconnect.CreateDirectConnectGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreateDirectConnectGatewayAssociationResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreateDirectConnectGatewayAssociationResult) Get(ctx workflow.Context) (*directconnect.CreateDirectConnectGatewayAssociationOutput, error) {
    var output directconnect.CreateDirectConnectGatewayAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreateDirectConnectGatewayAssociationProposalResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreateDirectConnectGatewayAssociationProposalResult) Get(ctx workflow.Context) (*directconnect.CreateDirectConnectGatewayAssociationProposalOutput, error) {
    var output directconnect.CreateDirectConnectGatewayAssociationProposalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreateInterconnectResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreateInterconnectResult) Get(ctx workflow.Context) (*directconnect.Interconnect, error) {
    var output directconnect.Interconnect
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreateLagResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreateLagResult) Get(ctx workflow.Context) (*directconnect.Lag, error) {
    var output directconnect.Lag
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreatePrivateVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreatePrivateVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreatePublicVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreatePublicVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectCreateTransitVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectCreateTransitVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.CreateTransitVirtualInterfaceOutput, error) {
    var output directconnect.CreateTransitVirtualInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDeleteBGPPeerResult struct {
	Result workflow.Future
}

func (r *DirectconnectDeleteBGPPeerResult) Get(ctx workflow.Context) (*directconnect.DeleteBGPPeerOutput, error) {
    var output directconnect.DeleteBGPPeerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDeleteConnectionResult struct {
	Result workflow.Future
}

func (r *DirectconnectDeleteConnectionResult) Get(ctx workflow.Context) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDeleteDirectConnectGatewayResult struct {
	Result workflow.Future
}

func (r *DirectconnectDeleteDirectConnectGatewayResult) Get(ctx workflow.Context) (*directconnect.DeleteDirectConnectGatewayOutput, error) {
    var output directconnect.DeleteDirectConnectGatewayOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDeleteDirectConnectGatewayAssociationResult struct {
	Result workflow.Future
}

func (r *DirectconnectDeleteDirectConnectGatewayAssociationResult) Get(ctx workflow.Context) (*directconnect.DeleteDirectConnectGatewayAssociationOutput, error) {
    var output directconnect.DeleteDirectConnectGatewayAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDeleteDirectConnectGatewayAssociationProposalResult struct {
	Result workflow.Future
}

func (r *DirectconnectDeleteDirectConnectGatewayAssociationProposalResult) Get(ctx workflow.Context) (*directconnect.DeleteDirectConnectGatewayAssociationProposalOutput, error) {
    var output directconnect.DeleteDirectConnectGatewayAssociationProposalOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDeleteInterconnectResult struct {
	Result workflow.Future
}

func (r *DirectconnectDeleteInterconnectResult) Get(ctx workflow.Context) (*directconnect.DeleteInterconnectOutput, error) {
    var output directconnect.DeleteInterconnectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDeleteLagResult struct {
	Result workflow.Future
}

func (r *DirectconnectDeleteLagResult) Get(ctx workflow.Context) (*directconnect.Lag, error) {
    var output directconnect.Lag
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDeleteVirtualInterfaceResult struct {
	Result workflow.Future
}

func (r *DirectconnectDeleteVirtualInterfaceResult) Get(ctx workflow.Context) (*directconnect.DeleteVirtualInterfaceOutput, error) {
    var output directconnect.DeleteVirtualInterfaceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeConnectionLoaResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeConnectionLoaResult) Get(ctx workflow.Context) (*directconnect.DescribeConnectionLoaOutput, error) {
    var output directconnect.DescribeConnectionLoaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeConnectionsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeConnectionsResult) Get(ctx workflow.Context) (*directconnect.Connections, error) {
    var output directconnect.Connections
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeConnectionsOnInterconnectResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeConnectionsOnInterconnectResult) Get(ctx workflow.Context) (*directconnect.Connections, error) {
    var output directconnect.Connections
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeDirectConnectGatewayAssociationProposalsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeDirectConnectGatewayAssociationProposalsResult) Get(ctx workflow.Context) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
    var output directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeDirectConnectGatewayAssociationsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeDirectConnectGatewayAssociationsResult) Get(ctx workflow.Context) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
    var output directconnect.DescribeDirectConnectGatewayAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeDirectConnectGatewayAttachmentsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeDirectConnectGatewayAttachmentsResult) Get(ctx workflow.Context) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
    var output directconnect.DescribeDirectConnectGatewayAttachmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeDirectConnectGatewaysResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeDirectConnectGatewaysResult) Get(ctx workflow.Context) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
    var output directconnect.DescribeDirectConnectGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeHostedConnectionsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeHostedConnectionsResult) Get(ctx workflow.Context) (*directconnect.Connections, error) {
    var output directconnect.Connections
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeInterconnectLoaResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeInterconnectLoaResult) Get(ctx workflow.Context) (*directconnect.DescribeInterconnectLoaOutput, error) {
    var output directconnect.DescribeInterconnectLoaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeInterconnectsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeInterconnectsResult) Get(ctx workflow.Context) (*directconnect.DescribeInterconnectsOutput, error) {
    var output directconnect.DescribeInterconnectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeLagsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeLagsResult) Get(ctx workflow.Context) (*directconnect.DescribeLagsOutput, error) {
    var output directconnect.DescribeLagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeLoaResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeLoaResult) Get(ctx workflow.Context) (*directconnect.Loa, error) {
    var output directconnect.Loa
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeLocationsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeLocationsResult) Get(ctx workflow.Context) (*directconnect.DescribeLocationsOutput, error) {
    var output directconnect.DescribeLocationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeTagsResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeTagsResult) Get(ctx workflow.Context) (*directconnect.DescribeTagsOutput, error) {
    var output directconnect.DescribeTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeVirtualGatewaysResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeVirtualGatewaysResult) Get(ctx workflow.Context) (*directconnect.DescribeVirtualGatewaysOutput, error) {
    var output directconnect.DescribeVirtualGatewaysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDescribeVirtualInterfacesResult struct {
	Result workflow.Future
}

func (r *DirectconnectDescribeVirtualInterfacesResult) Get(ctx workflow.Context) (*directconnect.DescribeVirtualInterfacesOutput, error) {
    var output directconnect.DescribeVirtualInterfacesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectDisassociateConnectionFromLagResult struct {
	Result workflow.Future
}

func (r *DirectconnectDisassociateConnectionFromLagResult) Get(ctx workflow.Context) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectListVirtualInterfaceTestHistoryResult struct {
	Result workflow.Future
}

func (r *DirectconnectListVirtualInterfaceTestHistoryResult) Get(ctx workflow.Context) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
    var output directconnect.ListVirtualInterfaceTestHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectStartBgpFailoverTestResult struct {
	Result workflow.Future
}

func (r *DirectconnectStartBgpFailoverTestResult) Get(ctx workflow.Context) (*directconnect.StartBgpFailoverTestOutput, error) {
    var output directconnect.StartBgpFailoverTestOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectStopBgpFailoverTestResult struct {
	Result workflow.Future
}

func (r *DirectconnectStopBgpFailoverTestResult) Get(ctx workflow.Context) (*directconnect.StopBgpFailoverTestOutput, error) {
    var output directconnect.StopBgpFailoverTestOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectTagResourceResult struct {
	Result workflow.Future
}

func (r *DirectconnectTagResourceResult) Get(ctx workflow.Context) (*directconnect.TagResourceOutput, error) {
    var output directconnect.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectUntagResourceResult struct {
	Result workflow.Future
}

func (r *DirectconnectUntagResourceResult) Get(ctx workflow.Context) (*directconnect.UntagResourceOutput, error) {
    var output directconnect.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectUpdateDirectConnectGatewayAssociationResult struct {
	Result workflow.Future
}

func (r *DirectconnectUpdateDirectConnectGatewayAssociationResult) Get(ctx workflow.Context) (*directconnect.UpdateDirectConnectGatewayAssociationOutput, error) {
    var output directconnect.UpdateDirectConnectGatewayAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectUpdateLagResult struct {
	Result workflow.Future
}

func (r *DirectconnectUpdateLagResult) Get(ctx workflow.Context) (*directconnect.Lag, error) {
    var output directconnect.Lag
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectconnectUpdateVirtualInterfaceAttributesResult struct {
	Result workflow.Future
}

func (r *DirectconnectUpdateVirtualInterfaceAttributesResult) Get(ctx workflow.Context) (*directconnect.UpdateVirtualInterfaceAttributesOutput, error) {
    var output directconnect.UpdateVirtualInterfaceAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DirectConnectStub struct {
    activities awsactivities.DirectConnectActivities
}

func NewDirectConnectStub() DirectConnectClient {
    return &DirectConnectStub{}
}

func (a *DirectConnectStub) AcceptDirectConnectGatewayAssociationProposal(ctx workflow.Context, input *directconnect.AcceptDirectConnectGatewayAssociationProposalInput) (*directconnect.AcceptDirectConnectGatewayAssociationProposalOutput, error) {
    var output directconnect.AcceptDirectConnectGatewayAssociationProposalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AcceptDirectConnectGatewayAssociationProposal, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AcceptDirectConnectGatewayAssociationProposalAsync(ctx workflow.Context, input *directconnect.AcceptDirectConnectGatewayAssociationProposalInput) *DirectconnectAcceptDirectConnectGatewayAssociationProposalResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AcceptDirectConnectGatewayAssociationProposal, input)
    return &DirectconnectAcceptDirectConnectGatewayAssociationProposalResult{Result: future}
}

func (a *DirectConnectStub) AllocateConnectionOnInterconnect(ctx workflow.Context, input *directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := workflow.ExecuteActivity(ctx, a.activities.AllocateConnectionOnInterconnect, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AllocateConnectionOnInterconnectAsync(ctx workflow.Context, input *directconnect.AllocateConnectionOnInterconnectInput) *DirectconnectAllocateConnectionOnInterconnectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AllocateConnectionOnInterconnect, input)
    return &DirectconnectAllocateConnectionOnInterconnectResult{Result: future}
}

func (a *DirectConnectStub) AllocateHostedConnection(ctx workflow.Context, input *directconnect.AllocateHostedConnectionInput) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := workflow.ExecuteActivity(ctx, a.activities.AllocateHostedConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AllocateHostedConnectionAsync(ctx workflow.Context, input *directconnect.AllocateHostedConnectionInput) *DirectconnectAllocateHostedConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AllocateHostedConnection, input)
    return &DirectconnectAllocateHostedConnectionResult{Result: future}
}

func (a *DirectConnectStub) AllocatePrivateVirtualInterface(ctx workflow.Context, input *directconnect.AllocatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := workflow.ExecuteActivity(ctx, a.activities.AllocatePrivateVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AllocatePrivateVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.AllocatePrivateVirtualInterfaceInput) *DirectconnectAllocatePrivateVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AllocatePrivateVirtualInterface, input)
    return &DirectconnectAllocatePrivateVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) AllocatePublicVirtualInterface(ctx workflow.Context, input *directconnect.AllocatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := workflow.ExecuteActivity(ctx, a.activities.AllocatePublicVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AllocatePublicVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.AllocatePublicVirtualInterfaceInput) *DirectconnectAllocatePublicVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AllocatePublicVirtualInterface, input)
    return &DirectconnectAllocatePublicVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) AllocateTransitVirtualInterface(ctx workflow.Context, input *directconnect.AllocateTransitVirtualInterfaceInput) (*directconnect.AllocateTransitVirtualInterfaceOutput, error) {
    var output directconnect.AllocateTransitVirtualInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AllocateTransitVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AllocateTransitVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.AllocateTransitVirtualInterfaceInput) *DirectconnectAllocateTransitVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AllocateTransitVirtualInterface, input)
    return &DirectconnectAllocateTransitVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) AssociateConnectionWithLag(ctx workflow.Context, input *directconnect.AssociateConnectionWithLagInput) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateConnectionWithLag, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AssociateConnectionWithLagAsync(ctx workflow.Context, input *directconnect.AssociateConnectionWithLagInput) *DirectconnectAssociateConnectionWithLagResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateConnectionWithLag, input)
    return &DirectconnectAssociateConnectionWithLagResult{Result: future}
}

func (a *DirectConnectStub) AssociateHostedConnection(ctx workflow.Context, input *directconnect.AssociateHostedConnectionInput) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateHostedConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AssociateHostedConnectionAsync(ctx workflow.Context, input *directconnect.AssociateHostedConnectionInput) *DirectconnectAssociateHostedConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateHostedConnection, input)
    return &DirectconnectAssociateHostedConnectionResult{Result: future}
}

func (a *DirectConnectStub) AssociateVirtualInterface(ctx workflow.Context, input *directconnect.AssociateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) AssociateVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.AssociateVirtualInterfaceInput) *DirectconnectAssociateVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateVirtualInterface, input)
    return &DirectconnectAssociateVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) ConfirmConnection(ctx workflow.Context, input *directconnect.ConfirmConnectionInput) (*directconnect.ConfirmConnectionOutput, error) {
    var output directconnect.ConfirmConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) ConfirmConnectionAsync(ctx workflow.Context, input *directconnect.ConfirmConnectionInput) *DirectconnectConfirmConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ConfirmConnection, input)
    return &DirectconnectConfirmConnectionResult{Result: future}
}

func (a *DirectConnectStub) ConfirmPrivateVirtualInterface(ctx workflow.Context, input *directconnect.ConfirmPrivateVirtualInterfaceInput) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error) {
    var output directconnect.ConfirmPrivateVirtualInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmPrivateVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) ConfirmPrivateVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.ConfirmPrivateVirtualInterfaceInput) *DirectconnectConfirmPrivateVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ConfirmPrivateVirtualInterface, input)
    return &DirectconnectConfirmPrivateVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) ConfirmPublicVirtualInterface(ctx workflow.Context, input *directconnect.ConfirmPublicVirtualInterfaceInput) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error) {
    var output directconnect.ConfirmPublicVirtualInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmPublicVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) ConfirmPublicVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.ConfirmPublicVirtualInterfaceInput) *DirectconnectConfirmPublicVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ConfirmPublicVirtualInterface, input)
    return &DirectconnectConfirmPublicVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) ConfirmTransitVirtualInterface(ctx workflow.Context, input *directconnect.ConfirmTransitVirtualInterfaceInput) (*directconnect.ConfirmTransitVirtualInterfaceOutput, error) {
    var output directconnect.ConfirmTransitVirtualInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ConfirmTransitVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) ConfirmTransitVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.ConfirmTransitVirtualInterfaceInput) *DirectconnectConfirmTransitVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ConfirmTransitVirtualInterface, input)
    return &DirectconnectConfirmTransitVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) CreateBGPPeer(ctx workflow.Context, input *directconnect.CreateBGPPeerInput) (*directconnect.CreateBGPPeerOutput, error) {
    var output directconnect.CreateBGPPeerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBGPPeer, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreateBGPPeerAsync(ctx workflow.Context, input *directconnect.CreateBGPPeerInput) *DirectconnectCreateBGPPeerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBGPPeer, input)
    return &DirectconnectCreateBGPPeerResult{Result: future}
}

func (a *DirectConnectStub) CreateConnection(ctx workflow.Context, input *directconnect.CreateConnectionInput) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreateConnectionAsync(ctx workflow.Context, input *directconnect.CreateConnectionInput) *DirectconnectCreateConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateConnection, input)
    return &DirectconnectCreateConnectionResult{Result: future}
}

func (a *DirectConnectStub) CreateDirectConnectGateway(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayInput) (*directconnect.CreateDirectConnectGatewayOutput, error) {
    var output directconnect.CreateDirectConnectGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDirectConnectGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreateDirectConnectGatewayAsync(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayInput) *DirectconnectCreateDirectConnectGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDirectConnectGateway, input)
    return &DirectconnectCreateDirectConnectGatewayResult{Result: future}
}

func (a *DirectConnectStub) CreateDirectConnectGatewayAssociation(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayAssociationInput) (*directconnect.CreateDirectConnectGatewayAssociationOutput, error) {
    var output directconnect.CreateDirectConnectGatewayAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDirectConnectGatewayAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreateDirectConnectGatewayAssociationAsync(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayAssociationInput) *DirectconnectCreateDirectConnectGatewayAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDirectConnectGatewayAssociation, input)
    return &DirectconnectCreateDirectConnectGatewayAssociationResult{Result: future}
}

func (a *DirectConnectStub) CreateDirectConnectGatewayAssociationProposal(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayAssociationProposalInput) (*directconnect.CreateDirectConnectGatewayAssociationProposalOutput, error) {
    var output directconnect.CreateDirectConnectGatewayAssociationProposalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDirectConnectGatewayAssociationProposal, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreateDirectConnectGatewayAssociationProposalAsync(ctx workflow.Context, input *directconnect.CreateDirectConnectGatewayAssociationProposalInput) *DirectconnectCreateDirectConnectGatewayAssociationProposalResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDirectConnectGatewayAssociationProposal, input)
    return &DirectconnectCreateDirectConnectGatewayAssociationProposalResult{Result: future}
}

func (a *DirectConnectStub) CreateInterconnect(ctx workflow.Context, input *directconnect.CreateInterconnectInput) (*directconnect.Interconnect, error) {
    var output directconnect.Interconnect
    err := workflow.ExecuteActivity(ctx, a.activities.CreateInterconnect, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreateInterconnectAsync(ctx workflow.Context, input *directconnect.CreateInterconnectInput) *DirectconnectCreateInterconnectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateInterconnect, input)
    return &DirectconnectCreateInterconnectResult{Result: future}
}

func (a *DirectConnectStub) CreateLag(ctx workflow.Context, input *directconnect.CreateLagInput) (*directconnect.Lag, error) {
    var output directconnect.Lag
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLag, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreateLagAsync(ctx workflow.Context, input *directconnect.CreateLagInput) *DirectconnectCreateLagResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateLag, input)
    return &DirectconnectCreateLagResult{Result: future}
}

func (a *DirectConnectStub) CreatePrivateVirtualInterface(ctx workflow.Context, input *directconnect.CreatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePrivateVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreatePrivateVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.CreatePrivateVirtualInterfaceInput) *DirectconnectCreatePrivateVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePrivateVirtualInterface, input)
    return &DirectconnectCreatePrivateVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) CreatePublicVirtualInterface(ctx workflow.Context, input *directconnect.CreatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
    var output directconnect.VirtualInterface
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePublicVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreatePublicVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.CreatePublicVirtualInterfaceInput) *DirectconnectCreatePublicVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePublicVirtualInterface, input)
    return &DirectconnectCreatePublicVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) CreateTransitVirtualInterface(ctx workflow.Context, input *directconnect.CreateTransitVirtualInterfaceInput) (*directconnect.CreateTransitVirtualInterfaceOutput, error) {
    var output directconnect.CreateTransitVirtualInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTransitVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) CreateTransitVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.CreateTransitVirtualInterfaceInput) *DirectconnectCreateTransitVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateTransitVirtualInterface, input)
    return &DirectconnectCreateTransitVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) DeleteBGPPeer(ctx workflow.Context, input *directconnect.DeleteBGPPeerInput) (*directconnect.DeleteBGPPeerOutput, error) {
    var output directconnect.DeleteBGPPeerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBGPPeer, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DeleteBGPPeerAsync(ctx workflow.Context, input *directconnect.DeleteBGPPeerInput) *DirectconnectDeleteBGPPeerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBGPPeer, input)
    return &DirectconnectDeleteBGPPeerResult{Result: future}
}

func (a *DirectConnectStub) DeleteConnection(ctx workflow.Context, input *directconnect.DeleteConnectionInput) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DeleteConnectionAsync(ctx workflow.Context, input *directconnect.DeleteConnectionInput) *DirectconnectDeleteConnectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConnection, input)
    return &DirectconnectDeleteConnectionResult{Result: future}
}

func (a *DirectConnectStub) DeleteDirectConnectGateway(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayInput) (*directconnect.DeleteDirectConnectGatewayOutput, error) {
    var output directconnect.DeleteDirectConnectGatewayOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectConnectGateway, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DeleteDirectConnectGatewayAsync(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayInput) *DirectconnectDeleteDirectConnectGatewayResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectConnectGateway, input)
    return &DirectconnectDeleteDirectConnectGatewayResult{Result: future}
}

func (a *DirectConnectStub) DeleteDirectConnectGatewayAssociation(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayAssociationInput) (*directconnect.DeleteDirectConnectGatewayAssociationOutput, error) {
    var output directconnect.DeleteDirectConnectGatewayAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectConnectGatewayAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DeleteDirectConnectGatewayAssociationAsync(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayAssociationInput) *DirectconnectDeleteDirectConnectGatewayAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectConnectGatewayAssociation, input)
    return &DirectconnectDeleteDirectConnectGatewayAssociationResult{Result: future}
}

func (a *DirectConnectStub) DeleteDirectConnectGatewayAssociationProposal(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayAssociationProposalInput) (*directconnect.DeleteDirectConnectGatewayAssociationProposalOutput, error) {
    var output directconnect.DeleteDirectConnectGatewayAssociationProposalOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectConnectGatewayAssociationProposal, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DeleteDirectConnectGatewayAssociationProposalAsync(ctx workflow.Context, input *directconnect.DeleteDirectConnectGatewayAssociationProposalInput) *DirectconnectDeleteDirectConnectGatewayAssociationProposalResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDirectConnectGatewayAssociationProposal, input)
    return &DirectconnectDeleteDirectConnectGatewayAssociationProposalResult{Result: future}
}

func (a *DirectConnectStub) DeleteInterconnect(ctx workflow.Context, input *directconnect.DeleteInterconnectInput) (*directconnect.DeleteInterconnectOutput, error) {
    var output directconnect.DeleteInterconnectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteInterconnect, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DeleteInterconnectAsync(ctx workflow.Context, input *directconnect.DeleteInterconnectInput) *DirectconnectDeleteInterconnectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteInterconnect, input)
    return &DirectconnectDeleteInterconnectResult{Result: future}
}

func (a *DirectConnectStub) DeleteLag(ctx workflow.Context, input *directconnect.DeleteLagInput) (*directconnect.Lag, error) {
    var output directconnect.Lag
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLag, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DeleteLagAsync(ctx workflow.Context, input *directconnect.DeleteLagInput) *DirectconnectDeleteLagResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteLag, input)
    return &DirectconnectDeleteLagResult{Result: future}
}

func (a *DirectConnectStub) DeleteVirtualInterface(ctx workflow.Context, input *directconnect.DeleteVirtualInterfaceInput) (*directconnect.DeleteVirtualInterfaceOutput, error) {
    var output directconnect.DeleteVirtualInterfaceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVirtualInterface, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DeleteVirtualInterfaceAsync(ctx workflow.Context, input *directconnect.DeleteVirtualInterfaceInput) *DirectconnectDeleteVirtualInterfaceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteVirtualInterface, input)
    return &DirectconnectDeleteVirtualInterfaceResult{Result: future}
}

func (a *DirectConnectStub) DescribeConnectionLoa(ctx workflow.Context, input *directconnect.DescribeConnectionLoaInput) (*directconnect.DescribeConnectionLoaOutput, error) {
    var output directconnect.DescribeConnectionLoaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConnectionLoa, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeConnectionLoaAsync(ctx workflow.Context, input *directconnect.DescribeConnectionLoaInput) *DirectconnectDescribeConnectionLoaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConnectionLoa, input)
    return &DirectconnectDescribeConnectionLoaResult{Result: future}
}

func (a *DirectConnectStub) DescribeConnections(ctx workflow.Context, input *directconnect.DescribeConnectionsInput) (*directconnect.Connections, error) {
    var output directconnect.Connections
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeConnectionsAsync(ctx workflow.Context, input *directconnect.DescribeConnectionsInput) *DirectconnectDescribeConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConnections, input)
    return &DirectconnectDescribeConnectionsResult{Result: future}
}

func (a *DirectConnectStub) DescribeConnectionsOnInterconnect(ctx workflow.Context, input *directconnect.DescribeConnectionsOnInterconnectInput) (*directconnect.Connections, error) {
    var output directconnect.Connections
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConnectionsOnInterconnect, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeConnectionsOnInterconnectAsync(ctx workflow.Context, input *directconnect.DescribeConnectionsOnInterconnectInput) *DirectconnectDescribeConnectionsOnInterconnectResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConnectionsOnInterconnect, input)
    return &DirectconnectDescribeConnectionsOnInterconnectResult{Result: future}
}

func (a *DirectConnectStub) DescribeDirectConnectGatewayAssociationProposals(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
    var output directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectConnectGatewayAssociationProposals, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeDirectConnectGatewayAssociationProposalsAsync(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) *DirectconnectDescribeDirectConnectGatewayAssociationProposalsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectConnectGatewayAssociationProposals, input)
    return &DirectconnectDescribeDirectConnectGatewayAssociationProposalsResult{Result: future}
}

func (a *DirectConnectStub) DescribeDirectConnectGatewayAssociations(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAssociationsInput) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
    var output directconnect.DescribeDirectConnectGatewayAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectConnectGatewayAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeDirectConnectGatewayAssociationsAsync(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAssociationsInput) *DirectconnectDescribeDirectConnectGatewayAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectConnectGatewayAssociations, input)
    return &DirectconnectDescribeDirectConnectGatewayAssociationsResult{Result: future}
}

func (a *DirectConnectStub) DescribeDirectConnectGatewayAttachments(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAttachmentsInput) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
    var output directconnect.DescribeDirectConnectGatewayAttachmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectConnectGatewayAttachments, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeDirectConnectGatewayAttachmentsAsync(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewayAttachmentsInput) *DirectconnectDescribeDirectConnectGatewayAttachmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectConnectGatewayAttachments, input)
    return &DirectconnectDescribeDirectConnectGatewayAttachmentsResult{Result: future}
}

func (a *DirectConnectStub) DescribeDirectConnectGateways(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewaysInput) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
    var output directconnect.DescribeDirectConnectGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectConnectGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeDirectConnectGatewaysAsync(ctx workflow.Context, input *directconnect.DescribeDirectConnectGatewaysInput) *DirectconnectDescribeDirectConnectGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeDirectConnectGateways, input)
    return &DirectconnectDescribeDirectConnectGatewaysResult{Result: future}
}

func (a *DirectConnectStub) DescribeHostedConnections(ctx workflow.Context, input *directconnect.DescribeHostedConnectionsInput) (*directconnect.Connections, error) {
    var output directconnect.Connections
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHostedConnections, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeHostedConnectionsAsync(ctx workflow.Context, input *directconnect.DescribeHostedConnectionsInput) *DirectconnectDescribeHostedConnectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeHostedConnections, input)
    return &DirectconnectDescribeHostedConnectionsResult{Result: future}
}

func (a *DirectConnectStub) DescribeInterconnectLoa(ctx workflow.Context, input *directconnect.DescribeInterconnectLoaInput) (*directconnect.DescribeInterconnectLoaOutput, error) {
    var output directconnect.DescribeInterconnectLoaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInterconnectLoa, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeInterconnectLoaAsync(ctx workflow.Context, input *directconnect.DescribeInterconnectLoaInput) *DirectconnectDescribeInterconnectLoaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInterconnectLoa, input)
    return &DirectconnectDescribeInterconnectLoaResult{Result: future}
}

func (a *DirectConnectStub) DescribeInterconnects(ctx workflow.Context, input *directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error) {
    var output directconnect.DescribeInterconnectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInterconnects, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeInterconnectsAsync(ctx workflow.Context, input *directconnect.DescribeInterconnectsInput) *DirectconnectDescribeInterconnectsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInterconnects, input)
    return &DirectconnectDescribeInterconnectsResult{Result: future}
}

func (a *DirectConnectStub) DescribeLags(ctx workflow.Context, input *directconnect.DescribeLagsInput) (*directconnect.DescribeLagsOutput, error) {
    var output directconnect.DescribeLagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLags, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeLagsAsync(ctx workflow.Context, input *directconnect.DescribeLagsInput) *DirectconnectDescribeLagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLags, input)
    return &DirectconnectDescribeLagsResult{Result: future}
}

func (a *DirectConnectStub) DescribeLoa(ctx workflow.Context, input *directconnect.DescribeLoaInput) (*directconnect.Loa, error) {
    var output directconnect.Loa
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLoa, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeLoaAsync(ctx workflow.Context, input *directconnect.DescribeLoaInput) *DirectconnectDescribeLoaResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLoa, input)
    return &DirectconnectDescribeLoaResult{Result: future}
}

func (a *DirectConnectStub) DescribeLocations(ctx workflow.Context, input *directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error) {
    var output directconnect.DescribeLocationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeLocations, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeLocationsAsync(ctx workflow.Context, input *directconnect.DescribeLocationsInput) *DirectconnectDescribeLocationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeLocations, input)
    return &DirectconnectDescribeLocationsResult{Result: future}
}

func (a *DirectConnectStub) DescribeTags(ctx workflow.Context, input *directconnect.DescribeTagsInput) (*directconnect.DescribeTagsOutput, error) {
    var output directconnect.DescribeTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeTagsAsync(ctx workflow.Context, input *directconnect.DescribeTagsInput) *DirectconnectDescribeTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input)
    return &DirectconnectDescribeTagsResult{Result: future}
}

func (a *DirectConnectStub) DescribeVirtualGateways(ctx workflow.Context, input *directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error) {
    var output directconnect.DescribeVirtualGatewaysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVirtualGateways, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeVirtualGatewaysAsync(ctx workflow.Context, input *directconnect.DescribeVirtualGatewaysInput) *DirectconnectDescribeVirtualGatewaysResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVirtualGateways, input)
    return &DirectconnectDescribeVirtualGatewaysResult{Result: future}
}

func (a *DirectConnectStub) DescribeVirtualInterfaces(ctx workflow.Context, input *directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error) {
    var output directconnect.DescribeVirtualInterfacesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeVirtualInterfaces, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DescribeVirtualInterfacesAsync(ctx workflow.Context, input *directconnect.DescribeVirtualInterfacesInput) *DirectconnectDescribeVirtualInterfacesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeVirtualInterfaces, input)
    return &DirectconnectDescribeVirtualInterfacesResult{Result: future}
}

func (a *DirectConnectStub) DisassociateConnectionFromLag(ctx workflow.Context, input *directconnect.DisassociateConnectionFromLagInput) (*directconnect.Connection, error) {
    var output directconnect.Connection
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateConnectionFromLag, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) DisassociateConnectionFromLagAsync(ctx workflow.Context, input *directconnect.DisassociateConnectionFromLagInput) *DirectconnectDisassociateConnectionFromLagResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateConnectionFromLag, input)
    return &DirectconnectDisassociateConnectionFromLagResult{Result: future}
}

func (a *DirectConnectStub) ListVirtualInterfaceTestHistory(ctx workflow.Context, input *directconnect.ListVirtualInterfaceTestHistoryInput) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
    var output directconnect.ListVirtualInterfaceTestHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVirtualInterfaceTestHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) ListVirtualInterfaceTestHistoryAsync(ctx workflow.Context, input *directconnect.ListVirtualInterfaceTestHistoryInput) *DirectconnectListVirtualInterfaceTestHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListVirtualInterfaceTestHistory, input)
    return &DirectconnectListVirtualInterfaceTestHistoryResult{Result: future}
}

func (a *DirectConnectStub) StartBgpFailoverTest(ctx workflow.Context, input *directconnect.StartBgpFailoverTestInput) (*directconnect.StartBgpFailoverTestOutput, error) {
    var output directconnect.StartBgpFailoverTestOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartBgpFailoverTest, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) StartBgpFailoverTestAsync(ctx workflow.Context, input *directconnect.StartBgpFailoverTestInput) *DirectconnectStartBgpFailoverTestResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartBgpFailoverTest, input)
    return &DirectconnectStartBgpFailoverTestResult{Result: future}
}

func (a *DirectConnectStub) StopBgpFailoverTest(ctx workflow.Context, input *directconnect.StopBgpFailoverTestInput) (*directconnect.StopBgpFailoverTestOutput, error) {
    var output directconnect.StopBgpFailoverTestOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopBgpFailoverTest, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) StopBgpFailoverTestAsync(ctx workflow.Context, input *directconnect.StopBgpFailoverTestInput) *DirectconnectStopBgpFailoverTestResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopBgpFailoverTest, input)
    return &DirectconnectStopBgpFailoverTestResult{Result: future}
}

func (a *DirectConnectStub) TagResource(ctx workflow.Context, input *directconnect.TagResourceInput) (*directconnect.TagResourceOutput, error) {
    var output directconnect.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) TagResourceAsync(ctx workflow.Context, input *directconnect.TagResourceInput) *DirectconnectTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &DirectconnectTagResourceResult{Result: future}
}

func (a *DirectConnectStub) UntagResource(ctx workflow.Context, input *directconnect.UntagResourceInput) (*directconnect.UntagResourceOutput, error) {
    var output directconnect.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) UntagResourceAsync(ctx workflow.Context, input *directconnect.UntagResourceInput) *DirectconnectUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &DirectconnectUntagResourceResult{Result: future}
}

func (a *DirectConnectStub) UpdateDirectConnectGatewayAssociation(ctx workflow.Context, input *directconnect.UpdateDirectConnectGatewayAssociationInput) (*directconnect.UpdateDirectConnectGatewayAssociationOutput, error) {
    var output directconnect.UpdateDirectConnectGatewayAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDirectConnectGatewayAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) UpdateDirectConnectGatewayAssociationAsync(ctx workflow.Context, input *directconnect.UpdateDirectConnectGatewayAssociationInput) *DirectconnectUpdateDirectConnectGatewayAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDirectConnectGatewayAssociation, input)
    return &DirectconnectUpdateDirectConnectGatewayAssociationResult{Result: future}
}

func (a *DirectConnectStub) UpdateLag(ctx workflow.Context, input *directconnect.UpdateLagInput) (*directconnect.Lag, error) {
    var output directconnect.Lag
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLag, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) UpdateLagAsync(ctx workflow.Context, input *directconnect.UpdateLagInput) *DirectconnectUpdateLagResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateLag, input)
    return &DirectconnectUpdateLagResult{Result: future}
}

func (a *DirectConnectStub) UpdateVirtualInterfaceAttributes(ctx workflow.Context, input *directconnect.UpdateVirtualInterfaceAttributesInput) (*directconnect.UpdateVirtualInterfaceAttributesOutput, error) {
    var output directconnect.UpdateVirtualInterfaceAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateVirtualInterfaceAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *DirectConnectStub) UpdateVirtualInterfaceAttributesAsync(ctx workflow.Context, input *directconnect.UpdateVirtualInterfaceAttributesInput) *DirectconnectUpdateVirtualInterfaceAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateVirtualInterfaceAttributes, input)
    return &DirectconnectUpdateVirtualInterfaceAttributesResult{Result: future}
}