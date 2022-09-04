import { faMinus, faPlus, faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";
import { Modal, Button, Form } from "react-bootstrap";
import { numberWithCommas } from "../utils/utils";

const ModalKeranjang = ({
    showModal,
    handleClose,
    keranjangDetail,
    jumlah,
    keterangan,
    tambah,
    kurang,
    changeHandler,
    handleSubmit,
    totalHarga,
    hapusPesanan
}) => {
    if (keranjangDetail) {
        return (
            <Modal show={showModal} onHide={handleClose}>
                <Form onSubmit={handleSubmit}>
                    <Modal.Header closeButton>
                        <Modal.Title>
                            {keranjangDetail.product.nama_product}{" "}
                            <strong>
                                (Rp. {numberWithCommas(keranjangDetail.product.harga_product)})
                            </strong>
                        </Modal.Title>
                    </Modal.Header>
                    <Modal.Body>

                        <Form.Group controlId="exampleForm.ControlInput1">
                            <Form.Label>Total Harga :</Form.Label>
                            <p>
                                <strong>
                                    Rp. {numberWithCommas(totalHarga)}
                                </strong>
                            </p>
                        </Form.Group>

                        <Form.Group controlId="exampleForm.ControlInput1">
                            <Form.Label>Jumlah :</Form.Label>
                            <br />
                            <Button variant="primary" size="sm" className="me-2" onClick={() => kurang()}>
                                <FontAwesomeIcon icon={faMinus} />
                            </Button>

                            <strong>{jumlah}</strong>

                            <Button variant="primary" size="sm" className="ms-2" onClick={() => tambah()}>
                                <FontAwesomeIcon icon={faPlus} />
                            </Button>
                        </Form.Group>

                        <Form.Group controlId="exampleForm.ControlTextarea1">
                            <Form.Label>Keterangan :</Form.Label>
                            <Form.Control
                                as="textarea"
                                rows="3"
                                name="keterangan"
                                placeholder="Contoh : Pedes, Nasi Setengah"
                                value={keterangan}
                                onChange={(event) => changeHandler(event)}
                            />
                        </Form.Group>

                    </Modal.Body>
                    <Modal.Footer>
                        <Button className="me-2" variant="danger" onClick={() => hapusPesanan(keranjangDetail.id)}>
                            <FontAwesomeIcon icon={faTrash} /> Hapus Pesanan
                        </Button>
                        <Button variant="primary" type="submit">
                            Simpan
                        </Button>
                    </Modal.Footer>
                </Form>
            </Modal>
        );
    }
};

export default ModalKeranjang;
