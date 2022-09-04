import React, { Component } from 'react'
import { Badge, Card, Col, ListGroup, Row } from 'react-bootstrap'
import { numberWithCommas } from '../utils/utils'
import ModalKeranjang from './ModalKeranjang'
import TotalBayar from './TotalBayar'
import axios from 'axios'
import { API_URL } from '../utils/constants'
import swal from 'sweetalert'

export default class Hasil extends Component {
  constructor(props) {
    super(props)

    this.state = {
      showModal: false,
      keranjangDetail: false,
      jumlah: 0,
      keterangan: '',
      totalHarga: 0,
    }
  }

  handleShow = (menuKeranjang) => {
    this.setState({
      showModal: true,
      keranjangDetail: menuKeranjang,
      jumlah: menuKeranjang.jumlah,
      keterangan: menuKeranjang.keterangan,
      totalHarga: menuKeranjang.total_harga
    })
  }

  handleClose = () => {
    this.setState({
      showModal: false
    })
  }

  tambah = () => {
    this.setState({
      jumlah: this.state.jumlah + 1,
      totalHarga:
        this.state.keranjangDetail.product.harga_product * (this.state.jumlah + 1),
    });
  };

  kurang = () => {
    if (this.state.jumlah !== 1) {
      this.setState({
        jumlah: this.state.jumlah - 1,
        totalHarga:
          this.state.keranjangDetail.product.harga_product * (this.state.jumlah - 1),
      });
    }
  };

  changeHandler = (event) => {
    this.setState({
      keterangan: event.target.value,
    });
  };

  handleSubmit = (event) => {
    event.preventDefault();

    this.handleClose();

    const data = {
      jumlah: this.state.jumlah,
      keterangan: this.state.keterangan,
    };

    axios
      .put(API_URL + "cart/" + this.state.keranjangDetail.id_cart, data)
      .then((res) => {
        this.props.getListKeranjang();
        swal({
          title: "Update Pesanan!",
          text: "Sukses Update Pesanan",
          icon: "success",
          button: false,
          timer: 1500,
        });
      })
      .catch((error) => {
        console.log("Error yaa ", error);
      });
  };

  hapusPesanan = (id) => {
    this.handleClose();

    axios
      .delete(API_URL + "cart/" + id)
      .then((res) => {
        this.props.getListKeranjang()
        swal({
          title: "Hapus Pesanan!",
          text:
            "Sukses Hapus Pesanan " + this.state.keranjangDetail.product.nama_product,
          icon: "error",
          button: false,
          timer: 1500,
        });
      })
      .catch((error) => {
        console.log("Error yaa ", error);
      });
  };

  render() {
    const { keranjangs } = this.props

    return (
      <Col md={3} mt='2' className='mt-3'>
        <h4><strong>Hasil</strong></h4>
        <hr />
        {keranjangs.length !== 0 && (
          <Card className='overflow-auto hasil'>
            <ListGroup variant="flush">
              {keranjangs.map((menuKeranjang) => (
                <ListGroup.Item
                  key={menuKeranjang.id_cart}
                  onClick={() => this.handleShow(menuKeranjang)}
                  style={{ cursor: 'pointer' }} >
                  <Row>
                    <Col xs={2}>
                      <h4>
                        <Badge pill variant="success" >
                          {menuKeranjang.jumlah}
                        </Badge>
                      </h4>
                    </Col>
                    <Col>
                      <h5 >{menuKeranjang.product.nama_product}</h5>
                      <p>Rp. {numberWithCommas(menuKeranjang.product.harga_product)} </p>
                    </Col>
                    <Col>
                      <strong className='float-right' >Rp. {numberWithCommas(menuKeranjang.total_harga)}</strong>
                    </Col>
                  </Row>
                </ListGroup.Item>


              ))}
            </ListGroup>
          </Card>
        )}

        <ModalKeranjang
          handleClose={this.handleClose}
          {...this.state}
          tambah={this.tambah}
          kurang={this.kurang}
          changeHandler={this.changeHandler}
          handleSubmit={this.handleSubmit}
          hapusPesanan={this.hapusPesanan}
        />
        <TotalBayar keranjangs={keranjangs} {...this.props} />

      </Col>
    )
  }
}
