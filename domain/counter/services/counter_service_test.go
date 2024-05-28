package services

import (
	"context"
	"testing"

	geometryDomain "counter/domain/counter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGeometryRepository struct {
	mock.Mock
}

func (m *MockGeometryRepository) GetGeometricObjects(ctx context.Context) ([]geometryDomain.StlObject, error) {
	args := m.Called(ctx, nil)
	return nil, args.Error(0)
}

func (m *MockGeometryRepository) SaveGeometricObject(ctx context.Context, g geometryDomain.GeometricObject) error {
	args := m.Called(ctx, g)
	return args.Error(0)
}

func TestNewAnalysisService(t *testing.T) {
	mockRepo := new(MockGeometryRepository)
	service := NewAnalysisService(mockRepo)
	assert.Equal(t, mockRepo, service.(*AnalysisService).geometryRepository)
}

func TestSTLAnalysis(t *testing.T) {
	mockRepo := new(MockGeometryRepository)
	mockRepo.On("SaveGeometricObject", mock.Anything, mock.Anything).Return(nil)

	service := &AnalysisService{
		geometryRepository: mockRepo,
	}

	geometricFigures := &geometryDomain.GeometricObject{
		Triangles: []geometryDomain.Triangles{
			{
				Vertex: []geometryDomain.Vertex{
					{V1: 0, V2: 0, V3: 0},
					{V1: 1, V2: 0, V3: 0},
					{V1: 1, V2: 1, V3: 1},
				},
			},
			{
				Vertex: []geometryDomain.Vertex{
					{V1: 0, V2: 0, V3: 0},
					{V1: 0, V2: 1, V3: 1},
					{V1: 1, V2: 1, V3: 1},
				},
			},
		},
	}

	err := service.STLAnalysis(context.Background(), geometricFigures)

	assert.NoError(t, err)
	assert.Equal(t, int32(len(geometricFigures.Triangles)), geometricFigures.QntTriangles)
	assert.Equal(t, service.CalculateSurfaceArea(*geometricFigures), geometricFigures.SurfaceArea)
	mockRepo.AssertExpectations(t)
}

func TestCalculateSurfaceArea(t *testing.T) {
	service := &AnalysisService{}

	geometricFigures := &geometryDomain.GeometricObject{
		Triangles: []geometryDomain.Triangles{
			{
				Vertex: []geometryDomain.Vertex{
					{V1: 0, V2: 0, V3: 0},
					{V1: 1, V2: 0, V3: 0},
					{V1: 1, V2: 1, V3: 1},
				},
			},
			{
				Vertex: []geometryDomain.Vertex{
					{V1: 0, V2: 0, V3: 0},
					{V1: 0, V2: 1, V3: 1},
					{V1: 1, V2: 1, V3: 1},
				},
			},
		},
	}

	expectedArea := 1.4142135623730956
	actualArea := service.CalculateSurfaceArea(*geometricFigures)
	assert.Equal(t, expectedArea, actualArea)
}
