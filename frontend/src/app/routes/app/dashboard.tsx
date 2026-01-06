import { Spinner } from "@/components/ui/spinner";
import { useGetPayments } from "@/features/payments/api/get-payments";
import PaymentDataTable from "@/features/payments/components/table";

export default function Dashboard() {
  const payments = useGetPayments();
  if (payments.isLoading) {
    return (
      <div className="pt-8 w-full items-center justify-center flex">
        <Spinner className="size-10" />
      </div>
    );
  }

  if (payments.isError) {
    return (
      <div className="pt-8 w-full items-center justify-center flex">
        Something went wrong...
      </div>
    );
  }

  if (payments.data) {
    return <PaymentDataTable data={payments.data} />;
  }
}
