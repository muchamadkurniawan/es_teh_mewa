package service

import (
	"context"
	"database/sql"
	"es_teh_mewa/helperMain"
	"es_teh_mewa/pesanan/helper"
	"es_teh_mewa/pesanan/model/repository"
	"es_teh_mewa/pesanan/web"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type PesananServiceImpl struct {
	PesananRepository repository.PesananRepository
	db                *sql.DB
}

func NewPesananServie(repo repository.PesananRepository, DB *sql.DB) PesananService {
	return &PesananServiceImpl{
		PesananRepository: repo,
		db:                DB,
	}
}

func (service *PesananServiceImpl) GetIdProduk(ctx context.Context) []string {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	find := service.PesananRepository.GetIdProduk(ctx, tx)
	return find
}

func (service *PesananServiceImpl) GetProdukJualsAll(ctx context.Context) []web.ProdukJual {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	find := service.PesananRepository.GetProdukJualsAll(ctx, tx)
	return find
}

func (service *PesananServiceImpl) GetProdukJual(ctx context.Context, id int) web.ProdukJual {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	find := service.PesananRepository.GetProdukJual(ctx, tx, id)
	return find
}

func (service *PesananServiceImpl) FindById(ctx context.Context, id int) (web.PesananRequestUpdate, error) {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	find := service.PesananRepository.FindById(ctx, tx, id)
	return find, err
}

func (service *PesananServiceImpl) FindAll(ctx context.Context) ([]web.PesananResponseDateString, error) {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	all := service.PesananRepository.FindAll(ctx, tx)
	var pesanans []web.PesananResponseDateString
	for _, pesanan := range all {
		newPesanan := helper.PesananEntityToResponseString(pesanan)
		pesanans = append(pesanans, newPesanan)
	}
	return pesanans, nil
}

func (service *PesananServiceImpl) FindPesananDetail(ctx context.Context, id int) web.PesananDetailUpdate {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	pesanan := service.PesananRepository.FindById(ctx, tx, id)
	detail := service.PesananRepository.ShowAllDetailPesananId(ctx, tx, pesanan.Id)
	b := web.PesananDetailUpdate{
		Pesanan: pesanan,
		Detail:  detail,
	}
	return b
}

func (service *PesananServiceImpl) FindAllPesananDetail(ctx context.Context) []web.PesananDetail {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	pesanan := service.PesananRepository.FindAllPesananDetail(ctx, tx)
	var pesananDetail []web.PesananDetail
	for _, a := range pesanan {
		detail := service.PesananRepository.FindAllDetailPesananId(ctx, tx, a.Id, 2)
		b := web.PesananDetail{
			Pesanan: a,
			Detail:  detail,
		}
		pesananDetail = append(pesananDetail, b)
	}
	return pesananDetail
}

func (service *PesananServiceImpl) CreatePesanan(ctx context.Context, request web.PesananRequestDateString) int {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	id := service.PesananRepository.CreatePesanan(ctx, tx, request)
	return id
}

func (service *PesananServiceImpl) CreateDetail(ctx context.Context, request web.DetailRequest) error {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	err = service.PesananRepository.CreateDetail(ctx, tx, request)
	return err
}

func (service *PesananServiceImpl) UpdatePembayaran(ctx context.Context, id int, pembayaran bool) error {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	err = service.PesananRepository.UpdatePembayaran(ctx, tx, id, pembayaran)
	helperMain.PanicIfError(err)
	return err
}

func (service *PesananServiceImpl) UpdateRekap(ctx context.Context, id int, id_rekap int) error {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	err = service.PesananRepository.UpdateRekap(ctx, tx, id, id_rekap)
	helperMain.PanicIfError(err)
	return err
}

func (service *PesananServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.db.Begin()
	helperMain.PanicIfError(err)
	defer helperMain.ErrorTx(tx)
	err = service.PesananRepository.Delete(ctx, tx, id)
	helperMain.PanicIfError(err)
	return err
}

func (service *PesananServiceImpl) Cetak(ctx context.Context, id int) error {
	tx, err := service.db.Begin()
	m := pdf.NewMaroto(consts.Portrait, consts.A5)
	m.SetPageMargins(20, 10, 20)
	helper.PdfHeader(m)
	tableHeading := []string{"Keterangan", "Harga", "Total"}
	m.SetBackgroundColor(color.NewWhite())
	data := service.PesananRepository.ShowAllDetailPesananId(ctx, tx, id)
	converData := helper.SetData(data)
	m.TableList(tableHeading, converData, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{5, 2, 2},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{5, 2, 2},
		},

		Align:              consts.Left,
		HeaderContentSpace: 2,
		Line:               true,
	})
	err = m.OutputFileAndClose("storage/struk.pdf")
	if err != nil {
		return err
	}
	return nil
}
