import PaymentDataTable from "@/features/payments/components/table";

const data = [
  {
    id: 1,
    sender: "Axel",
    recipient: "Tilde",
    amount: 100.5,
    createdAt: new Date("2025-12-30"),
    updatedAt: new Date("2025-12-30"),
  },
  {
    id: 2,
    sender: "Bob",
    recipient: "Sara",
    amount: 2225.0,
    createdAt: new Date("2026-01-01"),
    updatedAt: new Date("2026-01-01"),
  },
];

export default function Dashboard() {
  return <PaymentDataTable data={data} />;
}
