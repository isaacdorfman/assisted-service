package hostcommands

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/host/hostutil"
	"github.com/openshift/assisted-service/models"
)

var _ = Describe("domainNameResolution", func() {
	ctx := context.Background()
	var host models.Host
	var cluster common.Cluster
	var db *gorm.DB
	var dCmd *domainNameResolutionCmd
	var id, clusterID, infraEnvID strfmt.UUID
	var stepReply []*models.Step
	var stepErr error
	var dbName string
	var name string
	var baseDNSDomain string

	BeforeEach(func() {
		db, dbName = common.PrepareTestDB()
		dCmd = NewDomainNameResolutionCmd(common.GetTestLog(), "quay.io/ocpmetal/assisted-installer-agent:latest", db)
		id = strfmt.UUID(uuid.New().String())
		clusterID = strfmt.UUID(uuid.New().String())
		infraEnvID = strfmt.UUID(uuid.New().String())
		host = hostutil.GenerateTestHost(id, infraEnvID, clusterID, models.HostStatusPreparingForInstallation)
		host.Inventory = hostutil.GenerateMasterInventory()
		Expect(db.Create(&host).Error).ShouldNot(HaveOccurred())
		name = "example"
		baseDNSDomain = "test.com"
	})

	It("happy flow", func() {
		cluster = common.Cluster{Cluster: models.Cluster{ID: &clusterID, Name: name, BaseDNSDomain: baseDNSDomain}}
		Expect(db.Create(&cluster).Error).ShouldNot(HaveOccurred())
		stepReply, stepErr = dCmd.GetSteps(ctx, &host)
		Expect(stepReply).ToNot(BeNil())
		Expect(stepReply[0].StepType).To(Equal(models.StepTypeDomainResolution))
		Expect(stepErr).ShouldNot(HaveOccurred())
	})

	It("Missing cluster name", func() {
		cluster = common.Cluster{Cluster: models.Cluster{ID: &clusterID, BaseDNSDomain: baseDNSDomain}}
		Expect(db.Create(&cluster).Error).ShouldNot(HaveOccurred())
		stepReply, stepErr = dCmd.GetSteps(ctx, &host)
		Expect(stepReply).To(BeNil())
		Expect(stepErr).To(HaveOccurred())
	})

	It("Missing domain base name", func() {
		cluster = common.Cluster{Cluster: models.Cluster{ID: &clusterID, Name: name}}
		Expect(db.Create(&cluster).Error).ShouldNot(HaveOccurred())
		stepReply, stepErr = dCmd.GetSteps(ctx, &host)
		Expect(stepReply).To(BeNil())
		Expect(stepErr).To(HaveOccurred())
	})

	AfterEach(func() {
		common.DeleteTestDB(db, dbName)
		stepReply = nil
		stepErr = nil
	})
})
