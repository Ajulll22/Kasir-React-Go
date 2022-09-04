import React, { Component } from 'react'
import { Col, Container, Row } from 'react-bootstrap';
import { API_URL } from '../utils/constants';
import axios from 'axios'
import swal from 'sweetalert'
import { Hasil, Menus, ListCategories } from '../components'



export default class Home extends Component {
    constructor(props) {
        super(props)

        this.state = {
            menus: [],
            dipilih: "",
            keranjangs: []
        }
    }

    componentDidMount() {
        axios.get(API_URL + 'products?id_category=' + this.state.dipilih)
            .then(res => {
                const menus = res.data;
                this.setState({ menus });
            })
            .catch(error => {
                console.log(error);
            })

        this.getListKeranjang();
    }

    // componentDidUpdate(prevState) {
    //     if (this.state.keranjangs !== prevState.keranjangs) {
    //         axios.get(API_URL + 'keranjangs')
    //             .then(res => {
    //                 const keranjangs = res.data;
    //                 this.setState({ keranjangs });
    //             })
    //             .catch(error => {
    //                 console.log(error);
    //             })
    //     }
    // }

    getListKeranjang = () => {
        axios.get(API_URL + 'cart')
            .then(res => {
                const keranjangs = res.data;
                this.setState({ keranjangs });
            })
            .catch(error => {
                console.log(error);
            })
    }

    changeCategory = (value) => {
        this.setState({
            dipilih: value,
            menu: []
        })
        console.log(value);

        axios.get(API_URL + 'products?id_category=' + value)
            .then(res => {
                const menus = res.data;
                this.setState({ menus });
            })
            .catch(error => {
                console.log(error);
            })
    }

    masukKeranjang = (value) => {
        const keranjang = {
            id_product: value
        }
        axios.post(API_URL + 'cart', keranjang)
            .then(res => {
                this.getListKeranjang()
                swal({
                    title: "Success",
                    text: "Berhasil Masuk Keranjang ",
                    icon: "success",
                    button: false,
                    timer: 1000,
                });
            })
            .catch(error => {
                console.log(error);
            })

    }


    render() {
        const { menus, dipilih, keranjangs } = this.state
        return (
            <div className='mt-3'>
                <Container fluid>

                    <Row>
                        <ListCategories changeCategory={this.changeCategory} dipilih={dipilih} />
                        <Col className='mt-3' >
                            <h4><strong>Daftar Produk</strong></h4>
                            <hr />
                            <Row className='overflow-auto menu'>
                                {menus && menus.map((menu) => (<Menus
                                    key={menu.id_product}
                                    menu={menu}
                                    masukKeranjang={this.masukKeranjang}
                                />))}
                            </Row>
                        </Col>
                        <Hasil keranjangs={keranjangs} {...this.props} getListKeranjang={this.getListKeranjang} />
                    </Row>
                </Container>
            </div>
        )
    }
}

