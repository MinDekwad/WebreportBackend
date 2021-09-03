package handler

import (
	"encoding/csv"
	"fmt"
	"go-api-report2/model"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/jszwec/csvutil"
	"github.com/labstack/echo/v4"
)

func (h Handler) Upload(c echo.Context) error {

	var result struct {
		Total int `json:"total,omitempty"`
	}

	slug := c.Param("slug")
	file, err := c.FormFile("file")
	user := c.QueryParam("user")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		h.logger.Errorln("[Handler] open file error", err)
		return err
	}
	defer src.Close()

	csvReader := csv.NewReader(src)
	csvReader.FieldsPerRecord = -1
	csvReader.Comma = ','
	csvReader.LazyQuotes = true

	if slug == "userprofileamlocsv" {
		csvReader.Comma = '\t'
	}
	if slug == "watchlistcsv" {
		csvReader.Comma = '\t'
		csvReader.Comma = ','
	}

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		h.logger.Errorln("[Handler] new decode csv err", err)
		fmt.Println("err : ", err)
		return err
	}

	dec.Register(model.UnmarshalTime)

	var count int
	var data []interface{}
	filename := file.Filename

	switch slug {
	case "walletcsv":
		// data = make([]model.WalletDaily, 0)
		for {
			var w model.WalletDaily
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}

	case "consumercsv":
		for {
			var w model.Consumer
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}

	case "merchantcsv":
		for {
			var w model.Merchant
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}

	case "agentkyccsv":
		for {
			var w model.AgentKycCSV
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}

	case "loanbindingcsv":
		for {
			var w model.LoanbindingCSV
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}

	case "kycpendingcsv":
		for {
			var w model.Kycpending
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}

	case "lbpendingcsv":
		for {
			var w model.Lbpending
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}

	// new
	case "userprofileamlocsv":
		dec.Register(func(data []byte, n *int) error {
			var value int
			s := string(data)
			if strings.Contains(s, "NULL") {
				*n = value
				return nil
			}

			value, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			*n = value

			return nil
		})

		for {
			var w model.UserProfileAmlo
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				h.logger.Errorln("[Handler] dec Decode error", err)
				return err
			}
			data = append(data, w)
		}

		go func() {
			err = h.reportService.PostDataUpload(data, len(data), slug, filename, user)
			if err != nil {
				h.logger.Errorf("[Handler] Report %s insert data error", err)
			}
		}()
		result.Total = len(data)
		return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})

	case "watchlistcsv":
		for {
			var w model.WathclistCSVImport
			if err := dec.Decode(&w); err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			data = append(data, w)
		}

	}

	count = len(data)
	err = h.reportService.PostDataUpload(data, count, slug, filename, user)
	if err != nil {
		h.logger.Errorf("[Handler] Report %s insert data error", err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500107,
			Message: "Report insert data error",
		})
	}

	return c.JSON(http.StatusOK, model.Response{Message: "success", Data: result})
}
