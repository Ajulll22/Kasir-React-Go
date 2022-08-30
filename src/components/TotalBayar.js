import { faShoppingCart } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React, { Component } from 'react'
import { Button, Col, Row } from 'react-bootstrap'
import { numberWithCommas } from '../utils/utils';

export default class TotalBayar extends Component {
    render() {
        const totalBayar = this.props.keranjangs.reduce(function (result, item) {
            return result + item.total_harga;
        }, 0);

        return (
            <div className='fixed-bottom'>
                <Row>
                    <Col md={{ span: 3, offset: 9 }} className="px-4">
                        <h4>
                            Total Harga : <strong className="float-end ms-2"> Rp. {numberWithCommas(totalBayar)} </strong>
                        </h4>

                        <div className="d-grid gap-2 mb-3 mt-4">
                            <Button variant="primary" size="lg">
                                <FontAwesomeIcon icon={faShoppingCart} /> <strong>Bayar</strong>
                            </Button>
                        </div>
                    </Col>
                </Row>
            </div>
        )
    }
}
