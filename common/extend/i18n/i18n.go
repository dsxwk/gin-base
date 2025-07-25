package i18n

import (
	"embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	//go:embed locales/**/*.json
	embedFS        embed.FS
	Bundle         *i18n.Bundle
	supportedLangs = []string{"zh", "en"}
	Localizers     = LocalizerMap{} // 全局翻译集合
)

type LocalizerMap map[string]*i18n.Localizer

// Init 初始化翻译资源和 Localizers
func Init(useEmbed bool) {
	Bundle = i18n.NewBundle(language.Chinese)
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	Bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var err error
	if useEmbed {
		err = loadAllFromEmbed()
	} else {
		err = loadAllFromDisk()
	}
	if err != nil {
		log.Fatalf("i18n init error: %v", err)
	}

	// 初始化全局 localizers
	for _, lang := range supportedLangs {
		Localizers[lang] = i18n.NewLocalizer(Bundle, lang)
	}
}

// 递归加载 embed 下所有 json 文件
func loadAllFromEmbed() error {
	return fs.WalkDir(embedFS, "locales", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || (!strings.HasSuffix(path, ".json") && !strings.HasSuffix(path, ".yaml")) {
			return nil
		}

		data, err := embedFS.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = Bundle.ParseMessageFileBytes(data, path)

		return err
	})
}

// 递归加载磁盘目录
func loadAllFromDisk() error {
	return filepath.Walk("common/extend/i18n/locales", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || (!strings.HasSuffix(path, ".json") && !strings.HasSuffix(path, ".yaml")) {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = Bundle.ParseMessageFileBytes(data, path)

		return err
	})
}

// T 获取翻译
func T(c *gin.Context, id string, data map[string]interface{}) string {
	str, _ := c.Get("lang")
	lang, ok := str.(string)
	if !ok || str == "" {
		lang = "zh"
	}

	l, ok := Localizers[lang]
	if !ok {
		l = Localizers["zh"]
	}
	msg, err := l.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
	})
	if err != nil {
		return id
	}

	return msg
}
