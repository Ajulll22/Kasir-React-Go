import React from 'react'
import { Card, Col } from 'react-bootstrap'
import { numberWithCommas } from '../utils/utils'

const Menus = ({ menu, masukKeranjang }) => {
    return (
        <Col md={4} xs={6} className="mb-4" >
            <Card className='shadow' onClick={() => masukKeranjang(menu.id_product)} >
                <Card.Img style={{ height: '14rem' }} variant="top" src={"assets/images/" + menu.category.nama_category.toLowerCase() + "/" + menu.gambar_product} />
                <Card.Body>
                    <Card.Title>{menu.nama_product}</Card.Title>
                    <Card.Text>
                        Rp. {numberWithCommas(menu.harga_product)}
                    </Card.Text>
                </Card.Body>
            </Card>
        </Col>
    )
}

export default Menus