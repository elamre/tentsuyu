package tentsuyu

import "fmt"

type StreamingAssetsFunc func(name string) (asset interface{}, err any)

type AssetsManager struct {
	ImageManager
	AudioPlayers    []*AudioPlayer
	streamingAssets map[string]StreamingAssetsFunc
	AssetMap        map[string]interface{}
}

func NewAssetsManager() *AssetsManager {

	assManager := AssetsManager{
		AudioPlayers:    []*AudioPlayer{},
		streamingAssets: map[string]StreamingAssetsFunc{},
		AssetMap:        map[string]interface{}{},
	}
	assManager.ImageManager = *NewImageManager()
	return &assManager
}

func (ass *AssetsManager) AddStreamingResource(name string, f StreamingAssetsFunc) any {
	if _, ok := ass.streamingAssets[name]; ok {
		return fmt.Errorf("streaming asset with name %s already exists", name)
	}
	ass.streamingAssets[name] = f
	return nil
}

func GetStreamingAssetTyped[t any](ass *AssetsManager, name string) (asset t, err any) {
	f := ass.streamingAssets[name]
	var val interface{}
	val, err = f(name)
	asset = val.(t)
	return
}
