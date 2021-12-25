package readisearch

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"strconv"
)

type Redisearch struct {
	rs     *redisearch.Client
	logger *zap.Logger
}

func NewReadisearch(cfg Config, logger *zap.Logger) (*Redisearch, error) {
	addr:=fmt.Sprintf("%s:%d",cfg.Host,cfg.Port)
	c := redisearch.NewClient(addr, "myIndex")
	k,e:=c.Info()
	fmt.Println(k,e)
	s := &Redisearch{
		rs: c,
		logger: logger,
	}
	return s, nil
}

func (r *Redisearch) Migrate() error {
	sc:=redisearch.NewSchema(redisearch.DefaultOptions).AddField(redisearch.NewNumericField("product_id")).
		AddField(redisearch.NewTextFieldOptions("name",redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewTextField("brand")).
		AddField(redisearch.NewTextField("category")).
		AddField(redisearch.NewNumericField("quantity")).
		AddField(redisearch.NewNumericField("price"))

	err:=r.rs.Drop()
	fmt.Println(err)
	err=r.rs.CreateIndex(sc)
	if err!=nil{
		fmt.Printf("Error in create index %+v\n", err)
	}
	return err
}

//InitDB	This function insert test data to database
func (r *Redisearch) InitDB() {
	products:=[]Product{
		{1,"show","adidias","sport",2021,100,20},
		{2,"winter show ","adidias","sport",2021,100,20},
		{3,"summer show","adidias","sport",2021,100,20},
		{4,"bag","adidias","sport",2021,100,20},
		{5,"back bag","adidias","sport",2021,100,20},
		{6,"Tshirt","adidias","sport",2021,100,20},
		{7,"lap top","apple","digital",2021,100,20},
		{8,"lap top Mk200","apple","digital",2021,100,20},
		{9,"headphone","apple","digital",2021,100,20},
		{10,"headphone k963","apple","digital",2021,100,20},
	}
	var docs []redisearch.Document
	for _, p := range products {
		d:=redisearch.NewDocument(strconv.Itoa(p.ProductId),1)
		d.Set("product_id",p.ProductId).Set("name",p.Name).Set("brand",p.Brand).
			Set("category",p.Category).
			Set("model_yead",p.ModelYear).
			Set("price",p.Price).
			Set("quantity",p.Quantity)
		docs=append(docs,d)
	}
	err:=r.rs.Index(docs...)
	fmt.Println(err)
}
//
func (r *Redisearch) ProductSearch(term string) ([]Product, error) {
	var products []Product
	docs,_,err:=r.rs.Search(redisearch.NewQuery(term))

	if err!=nil{
		return nil, err
	}
	p:=Product{}
	for _, doc := range docs {
		doc.Properties["price"],_=strconv.ParseUint(doc.Properties["price"].(string),10,64)
		doc.Properties["quantity"],_=strconv.ParseUint(doc.Properties["quantity"].(string),10,64)
		mapstructure.Decode(doc.Properties,&p)
		products=append(products,p)
	}

	return products,nil
}

