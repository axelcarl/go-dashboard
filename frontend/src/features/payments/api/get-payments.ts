import { api } from "@/lib/api-client";
import { PaymentsApiSchema, type Payments } from "@/types/api";
import { useQuery } from "@tanstack/react-query";

async function getPayments(): Promise<Payments> {
  const response = await api.get("/payments");
  try {
    const payments = PaymentsApiSchema.parse(response);
    return payments;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export function useGetPayments() {
  return useQuery({
    queryKey: ["payments"],
    queryFn: getPayments,
  });
}
