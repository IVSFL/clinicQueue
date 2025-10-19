import React, { useEffect, useState } from "react";
import PatientInfo from "./PatientInfo";
import PatientQueue from "./PatientQueue";
import axios from "axios";

const DoctorAccountStyle = () => {
  const [currentPatient, setCurrentPatient] = useState(null);
  const [queue, setQueue] = useState([]);

  const doctorId = localStorage.getItem("id");

  async function fetchQueue() {
    try {
      const res = await axios.get(`http://localhost:8000/queue/${doctorId}`);
      setQueue(res.data);
    } catch (err) {
      console.error(err);
    }
  }

  useEffect(() => {
    fetchQueue();
  }, []);

  // const callPatient = (queueItem) => {
  //   setCurrentPatient({
  //     name: `${queueItem.patient.last_name} ${queueItem.patient.first_name} ${queueItem.patient.middle_name}`,
  //     ticketNumber: queueItem.ticket.ticket_number,
  //   });
  // };

  const callNextPatient = async () => {
    try {
      const res = await axios.post(
        `http://localhost:8000/queue/${doctorId}/call-next`
      );
      console.log(res.data)
      setCurrentPatient({
        name: `${res.data.patient.last_name} ${res.data.patient.first_name} ${res.data.patient.middle_name}`,
        ticketNumber: res.data.ticket_number,
      });
      fetchQueue()
    } catch (err) {
      console.log("ERR!", err);
    }
  };

  const callPatient = async (queueItem) => {
    try {
      const res = await axios.post(`http://localhost:8000/queue/${doctorId}/call/${queueItem.patient.id}`);
      setCurrentPatient({
      name: `${queueItem.patient.last_name} ${queueItem.patient.first_name} ${queueItem.patient.middle_name}`,
      content: queueItem.patient.content,
      ticketNumber: queueItem.ticket.ticket_number,
    });
    fetchQueue()
    } catch(err) {
      console.log("ERR! ", err)
    }
  }

  return (
    <div className="container py-4">
      <div className="row">
        <PatientInfo currentPatient={currentPatient} />
        <PatientQueue queue={queue} callPatient={callPatient} callNextPatient={callNextPatient}/>
      </div>
    </div>
  );
};

export default DoctorAccountStyle;
