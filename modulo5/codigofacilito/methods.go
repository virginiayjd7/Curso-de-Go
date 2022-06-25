package codigofacilito

func (self *Curso) getTitulo()string {//getTitulo es rivado  y publivo GetTitulo
	return self.Titulo
}
func (self *Curso) ObtenerTitulo()string {//getTitulo es rivado 
	return self.getTitulo()//llma al metodo privado
}