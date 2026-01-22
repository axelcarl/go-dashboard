import { api } from "@/lib/api-client";
import { type Payment } from "@/types/api";
import { useMutation } from "@tanstack/react-query";

async function createPayment(payment: Payment): Promise<void> {
  await api.post("/payments", {
    amount: payment.amount,
    sender: payment.sender,
    recipient: payment.recipient,
  });
}

export function useCreatePayment() {
  return useMutation({
    mutationFn: createPayment,
  });
}
