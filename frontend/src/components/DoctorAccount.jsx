import React, { useEffect, useState } from "react";
import PatientInfo from "./PatientInfo";
import PatientQueue from "./PatientQueue";
import DeferredPatients from "./DeferredPatient";
import OfficeModal from "./OfficeModal";
import TransferPatientModal from "./TransferPatientModal";
import axios from "axios";

const DoctorAccountStyle = () => {
  const [currentPatient, setCurrentPatient] = useState(null);
  const [queue, setQueue] = useState([]);
  const [deferredPatients, setDeferredPatients] = useState([]);
  const [doctorOffice, setDoctorOffice] = useState("");
  const [showTransferModal, setShowTransferModal] = useState(false);
  const [selectedTicketForTransfer, setSelectedTicketForTransfer] = useState("");
  const [showOfficeModal, setShowOfficeModal] = useState(false);

  const doctorId = localStorage.getItem("id");

  const fetchDoctorData = async () => {
    try {
      const res = await axios.get(`http://localhost:8000/doctors/${doctorId}`);
      setDoctorOffice(res.data.office || "");
    } catch (err) {
      console.error("Ошибка загрузки данных врача:", err);
    }
  };

  async function fetchQueue() {
    try {
      const res = await axios.get(`http://localhost:8000/queue/${doctorId}`);
      setQueue(res.data);
    } catch (err) {
      console.error(err);
    }
  }

  async function fetchDeferredPatients() {
    try {
      const res = await axios.get(
        `http://localhost:8000/queue/${doctorId}/deferred`
      );
      setDeferredPatients(res.data.deferred || []);
    } catch (err) {
      console.error("Ошибка загрузки отложенных пациентов:", err);
      setDeferredPatients([]);
    }
  }

  useEffect(() => {
    fetchDoctorData();
    fetchQueue();
    fetchDeferredPatients();
  }, []);

  const handleOfficeSet = (office) => {
    setDoctorOffice(office);
    setShowOfficeModal(false);
    fetchQueue();
    fetchDeferredPatients();
  };

  const callNextPatient = async () => {
    try {
      const res = await axios.post(
        `http://localhost:8000/queue/${doctorId}/call-next`
      );
      setCurrentPatient({
        name: `${res.data.patient.last_name} ${res.data.patient.first_name} ${res.data.patient.middle_name}`,
        content: res.data.patient.content,
        ticketNumber: res.data.ticket_number,
        office: res.data.office,
        patientId: res.data.patient.id,
      });
      fetchQueue();
      fetchDeferredPatients();
    } catch (err) {
      console.log("Ошибка вызова следующего пациента:", err);
    }
  };

  const callSpecificPatient = async (queueItem) => {
    try {
      const res = await axios.post(
        `http://localhost:8000/queue/${doctorId}/call/${queueItem.patient.id}`
      );
      setCurrentPatient({
        name: `${queueItem.patient.last_name} ${queueItem.patient.first_name} ${queueItem.patient.middle_name}`,
        content: queueItem.patient.content,
        ticketNumber: queueItem.ticket.ticket_number,
        office: res.data.office,
      });
      fetchQueue();
      fetchDeferredPatients();
    } catch (err) {
      console.log("Ошибка вызова пациента:", err);
    }
  };

  const callDeferredPatient = async (ticket) => {
    try {
      const res = await axios.post(
        `http://localhost:8000/queue/${doctorId}/call-deferred/${ticket.patient.id}`
      );
      setCurrentPatient({
        name: `${ticket.patient.last_name} ${ticket.patient.first_name} ${ticket.patient.middle_name}`,
        content: ticket.patient.content,
        ticketNumber: ticket.ticket_number,
        office: res.data.office,
      });
      fetchQueue();
      fetchDeferredPatients();
    } catch (err) {
      console.log("Ошибка вызова отложенного пациента:", err);
      alert("Не удалось вызвать отложенного пациента");
    }
  };

  const deferPatient = async (ticketNumber) => {
    try {
      await axios.post(`http://localhost:8000/queue/defer/${ticketNumber}`);
      fetchQueue();
      fetchDeferredPatients();
    } catch (err) {
      console.log("Ошибка откладывания пациента:", err);
    }
  };

  const completePatient = async (ticketNumber) => {
    try {
      await axios.post(`http://localhost:8000/queue/complete/${ticketNumber}`);
      setCurrentPatient(null);
      fetchQueue();
      fetchDeferredPatients();
    } catch (err) {
      console.log("Ошибка завершения приема:", err);
      alert("Не удалось завершить прием пациента");
    }
  };

  const handleTransferPatient = (ticketNumber) => {
    setSelectedTicketForTransfer(ticketNumber);
    setShowTransferModal(true);
  };

  const handleTransferClose = (shouldRefresh) => {
    setShowTransferModal(false);
    setSelectedTicketForTransfer("");
    
    if (shouldRefresh) {
      setCurrentPatient(null);
      fetchQueue();
      fetchDeferredPatients();
    }
  };

  return (
    <>
      <OfficeModal 
        doctorId={doctorId} 
        onOfficeSet={handleOfficeSet}
        show={showOfficeModal}
        onClose={() => setShowOfficeModal(false)}
      />
      
      <TransferPatientModal
        show={showTransferModal}
        onClose={handleTransferClose}
        ticketNumber={selectedTicketForTransfer}
        currentDoctorId={doctorId}
      />
      
      <div className="container-fluid py-4">
        {doctorOffice && (
          <div className="alert alert-info mb-3 d-flex justify-content-between align-items-center">
            <div>
              <strong>Текущий кабинет:</strong> {doctorOffice}
            </div>
            <button 
              className="btn btn-outline-secondary btn-sm"
              onClick={() => setShowOfficeModal(true)}
            >
              Сменить кабинет
            </button>
          </div>
        )}
        
        <div className="row">
          <div className="col-md-3">
            <div className="mb-4">
              <PatientInfo
                currentPatient={currentPatient}
                deferPatient={deferPatient}
                completePatient={completePatient}
                transferPatient={handleTransferPatient}
              />
            </div>
          </div>

          <div className="col-md-9">
            <PatientQueue
              queue={queue}
              callPatient={callSpecificPatient}
              callNextPatient={callNextPatient}
            />
          </div>
        </div>

        <div className="row">
          <div className="col-12">
            <DeferredPatients
              patients={deferredPatients}
              callPatient={callDeferredPatient}
            />
          </div>
        </div>
      </div>
    </>
  );
};

export default DoctorAccountStyle;